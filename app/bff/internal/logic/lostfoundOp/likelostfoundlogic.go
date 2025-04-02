package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLostFoundLogic {
	return &LikeLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLostFoundLogic) LikeLostFound(req *types.LikeLostFoundReq) (resp *types.LikeLostFoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
