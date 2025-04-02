package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteProductLogic {
	return &FavoriteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏/取消收藏商品
func (l *FavoriteProductLogic) FavoriteProduct(in *product.FavoriteProductRequest) (*product.FavoriteProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.FavoriteProductResponse{}, nil
}
