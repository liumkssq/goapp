package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProductsLogic {
	return &GetUserProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户的商品
func (l *GetUserProductsLogic) GetUserProducts(in *product.GetUserProductsRequest) (*product.GetUserProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.GetUserProductsResponse{}, nil
}
