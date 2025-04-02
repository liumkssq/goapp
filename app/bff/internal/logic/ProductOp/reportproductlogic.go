package ProductOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReportProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportProductLogic {
	return &ReportProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportProductLogic) ReportProduct(req *types.ReportProductReq) (resp *types.ReportProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
