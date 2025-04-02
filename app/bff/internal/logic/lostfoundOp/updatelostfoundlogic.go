package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLostFoundLogic {
	return &UpdateLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLostFoundLogic) UpdateLostFound(req *types.UpdateLostFoundReq) (resp *types.UpdateLostFoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
