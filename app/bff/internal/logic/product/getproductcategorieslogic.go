package product

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCategoriesLogic {
	return &GetProductCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductCategoriesLogic) GetProductCategories() (resp *types.ProductCategoriesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
