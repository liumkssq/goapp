package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/im/api/internal/svc"
	"github.com/liumkssq/goapp/app/im/api/internal/types"
	"github.com/liumkssq/goapp/app/im/rpc/imclient"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新会话
func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutConversationsLogic) PutConversations(req *types.PutConversationsReq) (resp *types.PutConversationsResp, err error) {
	uId := ctxdata.GetUidFromCtx(l.ctx)
	var conversationList map[string]*imclient.Conversation
	copier.Copy(&conversationList, req.ConversationList)
	_, err = l.svcCtx.PutConversations(l.ctx, &imclient.PutConversationsReq{
		UserId:           strconv.FormatInt(uId, 10),
		ConversationList: conversationList,
	})
	return
}
