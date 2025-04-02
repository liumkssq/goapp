package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLostFoundLogic {
	return &DeleteLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除失物招领
func (l *DeleteLostFoundLogic) DeleteLostFound(in *lostfound.DeleteLostFoundRequest) (*lostfound.DeleteLostFoundResponse, error) {
	// todo: add your logic here and delete this line

	return &lostfound.DeleteLostFoundResponse{}, nil
}
