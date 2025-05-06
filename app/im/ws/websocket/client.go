/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package websocket

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client interface {
	Close() error
	Send(v any) error
	SendUid(v any, uids ...string) error
	Read(v any) error
}

type client struct {
	*websocket.Conn
	host string

	opt dailOption

	Discover
}

func NewClient(host string, opts ...DailOptions) *client {
	opt := newDailOptions(opts...)

	c := client{
		Conn: nil,
		host: host,
		opt:  opt,
	}

	conn, err := c.dail()
	if err != nil {
		panic(err)
	}

	c.Conn = conn
	return &c
}

func (c *client) dail() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: c.host, Path: c.opt.pattern}

	// Log the connection attempt with headers for debugging
	fmt.Printf("Connecting to %s with headers: %v\n", u.String(), c.opt.header)

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	if err != nil {
		// Log response details for troubleshooting
		statusMsg := "no response"
		if resp != nil {
			statusMsg = fmt.Sprintf("HTTP %d", resp.StatusCode)
		}
		fmt.Printf("WebSocket connection failed: %v (Response: %s)\n", err, statusMsg)
		return nil, err
	}

	fmt.Printf("Successfully connected to WebSocket server: %s\n", u.String())
	return conn, err
}

func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Println("send data: ", string(data))

	err = c.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}
	// todo: 再增加一个重连发送

	// Ensure we use the same header and options for reconnection
	conn, err := c.dail()
	if err != nil {
		return err
	}
	c.Conn = conn
	return c.WriteMessage(websocket.TextMessage, data)
}

func (c *client) SendUid(v any, uids ...string) error {
	if c.Discover != nil {
		return c.Discover.Transpond(v, uids...)
	}
	return c.Send(v)
}

func (c *client) Read(v any) error {
	_, msg, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}

	return json.Unmarshal(msg, v)
}
