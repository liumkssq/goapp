package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/im/api/internal/svc"
	"github.com/liumkssq/goapp/app/im/api/internal/types"
	"github.com/liumkssq/goapp/app/im/rpc/imclient"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"strconv"

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

func (l *GetConversationsLogic) GetConversations(req *types.GetConversationsReq) (resp *types.GetConversationsResp, err error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println("uid", uid)
	data, err := l.svcCtx.GetConversations(l.ctx, &imclient.GetConversationsReq{
		UserId: strconv.FormatInt(uid, 10),
	})
	if err != nil {
		return nil, err
	}
	var res types.GetConversationsResp
	copier.Copy(&res, data)
	return &res, nil
}
