package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCommentsLogic {
	return &GetProductCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品评论
func (l *GetProductCommentsLogic) GetProductComments(in *product.GetProductCommentsRequest) (*product.GetProductCommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductCommentsResponse{}, nil
}
