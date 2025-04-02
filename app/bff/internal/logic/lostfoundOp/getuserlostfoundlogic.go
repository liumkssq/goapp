package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLostFoundLogic {
	return &GetUserLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLostFoundLogic) GetUserLostFound(req *types.UserLostFoundReq) (resp *types.UserLostFoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
