package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLostFoundLogic {
	return &UpdateLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新失物招领
func (l *UpdateLostFoundLogic) UpdateLostFound(in *lostfound.UpdateLostFoundRequest) (*lostfound.UpdateLostFoundResponse, error) {
	// todo: add your logic here and delete this line

	return &lostfound.UpdateLostFoundResponse{}, nil
}
