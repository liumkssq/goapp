package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteProductsLogic {
	return &GetFavoriteProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取收藏的商品
func (l *GetFavoriteProductsLogic) GetFavoriteProducts(in *product.GetFavoriteProductsRequest) (*product.GetFavoriteProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.GetFavoriteProductsResponse{}, nil
}
