package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkNotificationAsReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkNotificationAsReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkNotificationAsReadLogic {
	return &MarkNotificationAsReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 标记通知为已读
func (l *MarkNotificationAsReadLogic) MarkNotificationAsRead(in *user.MarkNotificationAsReadRequest) (*user.MarkNotificationAsReadResponse, error) {
	// todo: add your logic here and delete this line

	return &user.MarkNotificationAsReadResponse{}, nil
}
