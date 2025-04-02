package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLostFoundStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLostFoundStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLostFoundStatusLogic {
	return &UpdateLostFoundStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新失物招领状态
func (l *UpdateLostFoundStatusLogic) UpdateLostFoundStatus(in *lostfound.UpdateLostFoundStatusRequest) (*lostfound.UpdateLostFoundStatusResponse, error) {
	// todo: add your logic here and delete this line

	return &lostfound.UpdateLostFoundStatusResponse{}, nil
}
