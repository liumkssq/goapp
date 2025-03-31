package userOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkNotificationAsReadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 标记通知为已读
func NewMarkNotificationAsReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkNotificationAsReadLogic {
	return &MarkNotificationAsReadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarkNotificationAsReadLogic) MarkNotificationAsRead(req *types.MarkNotificationAsReadReq) (resp *types.MarkNotificationAsReadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
