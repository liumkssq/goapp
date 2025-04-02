package ProductOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteProductLogic {
	return &FavoriteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteProductLogic) FavoriteProduct(req *types.FavoriteProductReq) (resp *types.FavoriteProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
