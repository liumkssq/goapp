package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLostFoundStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLostFoundStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLostFoundStatusLogic {
	return &UpdateLostFoundStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLostFoundStatusLogic) UpdateLostFoundStatus(req *types.UpdateLostFoundStatusReq) (resp *types.UpdateLostFoundStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
