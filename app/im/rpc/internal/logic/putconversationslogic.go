package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/im/immodels"
	"github.com/liumkssq/goapp/app/im/rpc/im"
	"github.com/liumkssq/goapp/app/im/rpc/internal/svc"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/liumkssq/goapp/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新会话
func (l *PutConversationsLogic) PutConversations(in *im.PutConversationsReq) (*im.PutConversationsResp, error) {

	data, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("查询会话记录失败"), "ConversationsModel.FindByUserId err %v, req %v", err, in.UserId)
	}
	if data.ConversationList == nil {
		data.ConversationList = make(map[string]*immodels.Conversation)
	}
	for s, conversation := range in.ConversationList {
		var oldTotal int
		if data.ConversationList[s] != nil {
			oldTotal = data.ConversationList[s].Total
		}
		data.ConversationList[s] = &immodels.Conversation{
			ConversationId: conversation.ConversationId,
			ChatType:       constants.ChatType(conversation.ChatType),
			IsShow:         conversation.IsShow,
			Total:          int(conversation.Read) + oldTotal,
			Seq:            conversation.Seq,
		}
	}
	err = l.svcCtx.ConversationsModel.Update(l.ctx, data)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("更新会话记录失败"), "ConversationsModel.Update err %v, req %v", err, data)
	}
	return &im.PutConversationsResp{}, nil
}
