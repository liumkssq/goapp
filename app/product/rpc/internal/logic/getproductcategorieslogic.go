package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCategoriesLogic {
	return &GetProductCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品分类
func (l *GetProductCategoriesLogic) GetProductCategories(in *product.GetProductCategoriesRequest) (*product.GetProductCategoriesResponse, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductCategoriesResponse{}, nil
}
