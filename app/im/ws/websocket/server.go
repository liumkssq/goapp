/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/threading"

	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type AckType int

const (
	NoAck AckType = iota
	OnlyAck
	RigorAck
)

func (t AckType) ToString() string {
	switch t {
	case OnlyAck:
		return "OnlyAck"
	case RigorAck:
		return "RigorAck"
	}

	return "NoAck"
}

type Server struct {
	sync.RWMutex

	*threading.TaskRunner

	opt            *serverOption
	authentication Authentication

	routes     map[string]HandlerFunc
	addr       string
	patten     string
	listenOn   string
	discover   Discover
	connToUser map[*Conn]string
	userToConn map[string]*Conn

	upgrader websocket.Upgrader
	logx.Logger
}

func NewServer(addr string, opts ...ServerOptions) *Server {
	opt := newServerOptions(opts...)

	s := &Server{
		routes: make(map[string]HandlerFunc),
		addr:   addr,
		patten: opt.patten,
		opt:    &opt,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},

		authentication: opt.Authentication,

		connToUser: make(map[*Conn]string),
		userToConn: make(map[string]*Conn),
		discover:   opt.discover,
		listenOn:   FigureOutListenOn(addr),
		Logger:     logx.WithContext(context.Background()),
		TaskRunner: threading.NewTaskRunner(opt.concurrency),
	}

	// 存在服务发现，采用分布式im通信的时候; 默认不做任何处理
	s.discover.Register(fmt.Sprintf("%s", s.listenOn))
	fmt.Println("server start on %v", s.listenOn)

	return s
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("server handler ws recover err %v", r)
		}
	}()

	conn := NewConn(s, w, r)
	if conn == nil {
		return
	}

	if !s.authentication.Auth(w, r) {
		//conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint("不具备访问权限")))
		s.Send(&Message{FrameType: FrameData, Data: fmt.Sprint("不具备访问权限")}, conn)
		conn.Close()
		return
	}

	// 记录连接
	s.addConn(conn, r)

	// 处理连接
	go s.handlerConn(conn)
}

// 根据连接对象执行任务处理
func (s *Server) handlerConn(conn *Conn) {

	uids := s.GetUsers(conn)
	conn.Uid = uids[0]

	// 如果存在服务发现则进行注册；默认不做任何处理
	s.discover.BoundUser(conn.Uid)
	// 处理任务
	go s.handlerWrite(conn)

	if s.isAck(nil) {
		go s.readAck(conn)
	}

	for {
		// 获取请求消息
		_, msg, err := conn.ReadMessage()
		fmt.Println("new msg ", string(msg), err)
		if err != nil {
			s.Errorf("websocket conn read message err %v", err)
			s.Close(conn)
			return
		}
		// 解析消息
		var message Message
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal err %v, msg %v", err, string(msg))
			continue
		}

		// 依据消息进行处理
		if s.isAck(&message) {
			s.Infof("conn message read ack msg %v", message)
			conn.appendMsgMq(&message)
		} else {
			conn.message <- &message
		}
	}
}

func (s *Server) isAck(message *Message) bool {
	if message == nil {
		return s.opt.ack != NoAck
	}
	return s.opt.ack != NoAck && message.FrameType != FrameNoAck && message.FrameType != FrameTranspond
}

// 读取消息的ack
func (s *Server) readAck(conn *Conn) {
	for {
		select {
		case <-conn.done:
			s.Infof("close message ack uid %v ", conn.Uid)
			return
		default:
		}

		// 从队列中读取新的消息
		conn.messageMu.Lock()
		if len(conn.readMessage) == 0 {
			conn.messageMu.Unlock()
			// 增加睡眠
			time.Sleep(100 * time.Microsecond)
			continue
		}

		// 读取第一条
		message := conn.readMessage[0]

		// 判断ack的方式
		switch s.opt.ack {
		case OnlyAck:
			// 直接给客户端回复
			s.Send(&Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq + 1,
			}, conn)
			// 进行业务处理
			// 把消息从队列中移除
			conn.readMessage = conn.readMessage[1:]
			conn.messageMu.Unlock()

			conn.message <- message
		case RigorAck:
			// 先回
			if message.AckSeq == 0 {
				// 还未确认
				conn.readMessage[0].AckSeq++
				conn.readMessage[0].ackTime = time.Now()
				s.Send(&Message{
					FrameType: FrameAck,
					Id:        message.Id,
					AckSeq:    message.AckSeq,
				}, conn)
				s.Infof("message ack RigorAck send mid %v, seq %v , time%v", message.Id, message.AckSeq,
					message.ackTime)
				conn.messageMu.Unlock()
				continue
			}

			// 再验证

			// 1. 客户端返回结果，再一次确认
			// 得到客户端的序号
			msgSeq := conn.readMessageSeq[message.Id]
			if msgSeq.AckSeq > message.AckSeq {
				// 确认
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				conn.message <- message
				s.Infof("message ack RigorAck success mid %v", message.Id)
				continue
			}

			// 2. 客户端没有确认，考虑是否超过了ack的确认时间
			val := s.opt.ackTimeout - time.Since(message.ackTime)
			if !message.ackTime.IsZero() && val <= 0 {
				//		2.2 超过结束确认
				delete(conn.readMessageSeq, message.Id)
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				continue
			}
			//		2.1 未超过，重新发送
			conn.messageMu.Unlock()
			s.Send(&Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq,
			}, conn)
			// 睡眠一定的时间
			time.Sleep(3 * time.Second)
		}
	}
}

// 任务的处理
func (s *Server) handlerWrite(conn *Conn) {
	for {
		select {
		case <-conn.done:
			// 连接关闭
			return
		case message := <-conn.message:
			switch message.FrameType {
			case FramePing:
				s.Send(&Message{FrameType: FramePing}, conn)
			case FrameData:
				// 根据请求的method分发路由并执行
				if handler, ok := s.routes[message.Method]; ok {
					handler(s, conn, message)
				} else {
					s.Send(&Message{FrameType: FrameData, Data: fmt.Sprintf("不存在执行的方法 %v 请检查", message.Method)}, conn)
					//conn.WriteMessage(&Message{}, []byte(fmt.Sprintf("不存在执行的方法 %v 请检查", message.Method)))
				}
			}

			if s.isAck(message) {
				conn.messageMu.Lock()
				delete(conn.readMessageSeq, message.Id)
				conn.messageMu.Unlock()
			}
		}
	}
}

func (s *Server) addConn(conn *Conn, req *http.Request) {
	uid := s.authentication.UserId(req)
	s.Infof("[Debug][addConn] Attempting for UID: %s (conn: %p)", uid, conn) // 添加日志

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	if c := s.userToConn[uid]; c != nil {
		s.Info("[Debug][addConn] Overwriting existing entry for UID: %s. Old conn: %p, New conn: %p", uid, c, conn) // 修改日志
		// 暂时注释掉旧连接处理，避免复杂交互
		// delete(s.connToUser, c)
		// c.Close()
	}

	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
	keys := make([]string, 0, len(s.userToConn))
	for k := range s.userToConn {
		keys = append(keys, k)
	}
	s.Infof("[Debug][addConn] Successfully added/updated UID: %s (conn: %p). Current userToConn keys: %v", uid, conn, keys) // 添加日志
}

func (s *Server) GetConn(uid string) *Conn {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	fmt.Println(s.userToConn)
	return s.userToConn[uid]
}

func (s *Server) GetConns(uids ...string) []*Conn {
	if len(uids) == 0 {
		return nil
	}

	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	res := make([]*Conn, 0, len(uids))
	for _, uid := range uids {
		res = append(res, s.userToConn[uid])
	}
	return res
}

func (s *Server) GetUsers(conns ...*Conn) []string {

	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}

func (s *Server) Close(conn *Conn) {
	s.RWMutex.Lock() // 获取写锁

	uid := s.connToUser[conn]
	if uid == "" {
		s.Info("[Debug][Close] Attempted on already closed or unknown conn: %p", conn)
		s.RWMutex.Unlock() // 解锁
		return
	}

	s.Infof("[Debug][Close] Closing conn %p for UID: %s", conn, uid)
	delete(s.connToUser, conn)
	delete(s.userToConn, uid)

	keys := make([]string, 0, len(s.userToConn))
	for k := range s.userToConn {
		keys = append(keys, k)
	}
	s.RWMutex.Unlock() // 在调用 conn.Close() 前解锁

	conn.Close() // 实际关闭连接，这个调用不应该再持有 Server 的锁
	s.Infof("[Debug][Close] Conn %p closed for UID: %s. Remaining userToConn keys: %v", uid, keys)
}

func (s *Server) SendByUserId(msg interface{}, sendIds ...string) error {
	if len(sendIds) == 0 {
		return nil
	}

	return s.Send(msg, s.GetConns(sendIds...)...)
}

func (s *Server) Send(msg interface{}, conns ...*Conn) error {
	if len(conns) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, conn := range conns {
		if conn == nil {
			s.Infof("[Debug] Attempted to send message to a nil connection.")
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			s.Errorf("[Debug] Failed to write message to UID %s (conn: %p): %v", s.connToUser[conn], conn, err)
			s.Close(conn)
		} else {
			s.Infof("[Debug] Successfully sent message to UID %s (conn: %p)", s.connToUser[conn], conn)
		}
	}

	return nil
}

func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) Start() {
	http.HandleFunc(s.patten, s.ServerWs)
	s.Info(http.ListenAndServe(s.addr, nil))
}

func (s *Server) Stop() {
	fmt.Println("停止服务")
}
