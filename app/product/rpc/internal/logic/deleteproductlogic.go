package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商品
func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.DeleteProductResponse{}, nil
}
