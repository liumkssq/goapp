package immodels

import (
	"context"
	"fmt"
	"github.com/liumkssq/goapp/pkg/constants"
	"log"
	"sort"
	"strconv"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CombineId 函数，正确组合会话ID
func CombineId(aid, bid string) string {
	ids := []string{aid, bid}

	sort.Slice(ids, func(i, j int) bool {
		a, _ := strconv.ParseUint(ids[i], 0, 64)
		b, _ := strconv.ParseUint(ids[j], 0, 64)
		return a < b
	})

	return fmt.Sprintf("%s_%s", ids[0], ids[1])
}

func TestName(t *testing.T) {
	// 连接MongoDB
	ctx := context.Background()
	//mongodb://root:easy-chat@127.0.0.1:47017

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:easy-chat@localhost:27017"))

	if err != nil {
		log.Fatalf("连接MongoDB失败: %v", err)
	}
	defer client.Disconnect(ctx)

	// 选择数据库和集合
	db := client.Database("easy-chat") // 修改为您的数据库名
	collection := db.Collection("conversations")

	// 生成测试数据
	if err := generateConversationsData(ctx, collection); err != nil {
		log.Fatalf("生成测试数据失败: %v", err)
	}

	log.Println("测试数据生成成功!")
}

func generateConversationsData(ctx context.Context, collection *mongo.Collection) error {
	// 测试用户
	users := []string{"23", "21", "22", "20", "20"}

	now := time.Now()

	for _, userId := range users {
		// 创建会话列表
		conversations := &Conversations{
			ID:               primitive.NewObjectID(),
			UserId:           userId,
			ConversationList: make(map[string]*Conversation),
			CreateAt:         now,
			UpdateAt:         now,
		}

		// 生成3-5个会话
		numConversations := 3 + (len(userId) % 3)

		for i := 1; i <= numConversations; i++ {
			// 生成不同的目标用户ID
			targetUserId := fmt.Sprintf("%d", 1000+i)
			if targetUserId == userId {
				targetUserId = fmt.Sprintf("%d", 1001+i)
			}

			// 使用正确的方法创建会话ID
			conversationId := CombineId(userId, targetUserId)

			// 创建消息
			chatLog := &ChatLog{
				ID:             primitive.NewObjectID(),
				ConversationId: conversationId,
				SendId:         targetUserId,
				RecvId:         userId,
				MsgFrom:        1,                        // 假设1表示来自对方的消息
				ChatType:       constants.SingleChatType, // 假设是单聊
				MsgType:        constants.TextMType,      // 假设是文本消息
				MsgContent:     getRandomMessage(i),
				SendTime:       now.Add(-time.Duration(i*10) * time.Minute).Unix(),
				Status:         1, // 假设1表示已发送
				CreateAt:       now.Add(-time.Duration(i*10) * time.Minute),
				UpdateAt:       now.Add(-time.Duration(i*10) * time.Minute),
			}

			// 创建会话
			conversation := &Conversation{
				ID:             primitive.NewObjectID(),
				ConversationId: conversationId,
				ChatType:       constants.SingleChatType, // 单聊
				IsShow:         true,
				Total:          i * 5,                                       // 假设有一些消息历史
				Seq:            int64(i * 10),                               // 序列号
				Msg:            chatLog,                                     // 最后一条消息
				CreateAt:       now.Add(-time.Duration(i*60) * time.Minute), // 会话创建时间
				UpdateAt:       now.Add(-time.Duration(i*10) * time.Minute), // 最后更新时间与最后消息时间一致
			}

			// 添加到会话列表
			conversations.ConversationList[conversationId] = conversation
		}

		// 还可以添加一个群聊会话
		if len(userId)%2 == 0 {
			groupId := "group_" + userId[len(userId)-1:]
			groupConversationId := "group_" + groupId

			// 群聊消息
			groupChatLog := &ChatLog{
				ID:             primitive.NewObjectID(),
				ConversationId: groupConversationId,
				SendId:         "admin",                 // 可以是管理员或其他成员
				RecvId:         groupId,                 // 群ID作为接收者
				MsgFrom:        2,                       // 假设2表示群消息
				ChatType:       constants.GroupChatType, // 群聊
				MsgType:        constants.TextMType,
				MsgContent:     getGroupMessage(int(userId[len(userId)-1] - '0')),
				SendTime:       now.Add(-time.Duration(20) * time.Minute).Unix(),
				Status:         1,
				CreateAt:       now.Add(-time.Duration(20) * time.Minute),
				UpdateAt:       now.Add(-time.Duration(20) * time.Minute),
			}

			// 群聊会话
			groupConversation := &Conversation{
				ID:             primitive.NewObjectID(),
				ConversationId: groupConversationId,
				ChatType:       constants.GroupChatType, // 群聊
				IsShow:         true,
				Total:          20,  // 假设有20条历史消息
				Seq:            100, // 序列号
				Msg:            groupChatLog,
				CreateAt:       now.Add(-time.Duration(24*7) * time.Hour), // 一周前创建
				UpdateAt:       now.Add(-time.Duration(20) * time.Minute),
			}

			// 添加到会话列表
			conversations.ConversationList[groupConversationId] = groupConversation
		}

		// 插入数据库
		_, err := collection.InsertOne(ctx, conversations)
		if err != nil {
			return fmt.Errorf("为用户%s插入会话数据失败: %w", userId, err)
		}

		log.Printf("已为用户%s创建%d个会话", userId, len(conversations.ConversationList))
	}

	return nil
}

// 辅助函数
func getRandomMessage(seed int) string {
	messages := []string{
		"你好，有时间聊聊吗？",
		"下次考试的范围你知道吗？",
		"能把笔记发我一下吗？",
		"明天我们在图书馆见面",
		"这个问题我想了很久都没解决",
		"谢谢你的帮助！",
		"周末有什么安排？",
		"你的报告做完了吗？",
		"老师让我们明天交作业",
		"我刚刚发了个文件给你",
	}
	return messages[seed%len(messages)]
}

func getGroupMessage(seed int) string {
	messages := []string{
		"@全体成员 明天下午3点开会",
		"有谁知道这个问题的答案吗？",
		"我上传了一些学习资料到群文件",
		"请大家记得签到",
		"欢迎新成员加入",
	}
	return messages[seed%len(messages)]
}
