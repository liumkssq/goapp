package ProductOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCommentsLogic {
	return &GetProductCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductCommentsLogic) GetProductComments(req *types.ProductCommentsReq) (resp *types.ProductCommentsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
