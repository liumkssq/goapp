package ProductOp

import (
	"context"
	"errors"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"strconv"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	productClient "github.com/liumkssq/goapp/app/product/rpc/product"

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
	userId := ctxdata.GetUId(l.ctx)
	uid, _ := strconv.Atoi(userId)
	if uid == 0 {
		// Handle error: user not authenticated or userId not found in context
		// This should ideally be handled by middleware, but adding a check here is safer.
		return nil, errors.New("用户未登录或授权失败") // Consider using a standardized error type/code
	}

	rpcReq := productClient.ReportProductRequest{
		Id:          req.Id,
		UserId: int64(uid),
		Reason:      req.Reason,
		Description: req.Description,
		Images:      req.Images,
	}

	rpcResp, err := l.svcCtx.ProductRpc.ReportProduct(l.ctx, &rpcReq)
	if err != nil {
		l.Errorf("Failed to call ProductRpc.ReportProduct: %v", err)
		// Map the RPC error to a BFF error if needed
		return nil, err
	}

	resp = &types.ReportProductResp{
		ReportId: rpcResp.GetReportId(),
		Message:  rpcResp.GetMessage(),
	}

	return resp, nil
}
