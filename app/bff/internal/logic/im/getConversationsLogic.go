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

type GetConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取会话
func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConversationsLogic) GetConversations(
	req *types.GetConversationsReq) (resp *types.GetConversationsResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	data, err := l.svcCtx.ImRpc.GetConversations(l.ctx, &imclient.GetConversationsReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	var res types.GetConversationsResp
	res.UserId = uid
	copier.Copy(&res, data)
	return &res, nil
}
