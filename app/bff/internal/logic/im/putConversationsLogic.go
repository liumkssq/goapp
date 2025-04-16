package im

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/im/rpc/imclient"
	"github.com/liumkssq/goapp/pkg/ctxdata"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

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
	uId := ctxdata.GetUId(l.ctx)
	var conversationList map[string]*imclient.Conversation
	copier.Copy(&conversationList, req.ConversationList)
	_, err = l.svcCtx.ImRpc.PutConversations(l.ctx, &imclient.PutConversationsReq{
		UserId:           uId,
		ConversationList: conversationList,
	})
	return
}
