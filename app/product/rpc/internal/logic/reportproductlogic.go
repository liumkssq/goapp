package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReportProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportProductLogic {
	return &ReportProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 举报商品
func (l *ReportProductLogic) ReportProduct(in *product.ReportProductRequest) (*product.ReportProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ReportProductResponse{}, nil
}
