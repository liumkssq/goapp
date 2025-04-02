package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentProductLogic {
	return &CommentProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论商品
func (l *CommentProductLogic) CommentProduct(in *product.CommentProductRequest) (*product.CommentProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.CommentProductResponse{}, nil
}
