package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLostFoundCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLostFoundCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLostFoundCommentsLogic {
	return &GetLostFoundCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLostFoundCommentsLogic) GetLostFoundComments(req *types.LostFoundCommentsReq) (resp *types.LostFoundCommentsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
