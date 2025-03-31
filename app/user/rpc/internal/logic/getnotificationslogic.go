package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNotificationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNotificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNotificationsLogic {
	return &GetNotificationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取通知
func (l *GetNotificationsLogic) GetNotifications(in *user.GetNotificationsRequest) (*user.NotificationListResponse, error) {
	// todo: add your logic here and delete this line

	return &user.NotificationListResponse{}, nil
}
