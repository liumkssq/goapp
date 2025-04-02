package ProductOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProductsLogic {
	return &GetUserProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProductsLogic) GetUserProducts(req *types.UserProductsReq) (resp *types.UserProductsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
