package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLostFoundLogic {
	return &DeleteLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLostFoundLogic) DeleteLostFound(req *types.DeleteLostFoundReq) (resp *types.DeleteLostFoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
