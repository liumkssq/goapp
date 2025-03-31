package userOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNotificationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取通知
func NewGetNotificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNotificationsLogic {
	return &GetNotificationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNotificationsLogic) GetNotifications() (resp *types.NotificationListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
