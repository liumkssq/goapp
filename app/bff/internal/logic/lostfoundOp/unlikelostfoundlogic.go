package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnlikeLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeLostFoundLogic {
	return &UnlikeLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnlikeLostFoundLogic) UnlikeLostFound(req *types.UnlikeLostFoundReq) (resp *types.UnlikeLostFoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
