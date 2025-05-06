package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/liumkssq/goapp/app/product/model"
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
	var reportImages sql.NullString
	if len(in.GetImages()) > 0 {
		imagesJson, err := json.Marshal(in.GetImages())
		if err != nil {
			l.Errorf("Failed to marshal images to JSON: %v", err)
			return nil, err
		}
		reportImages = sql.NullString{String: string(imagesJson), Valid: true}
	}

	reportData := model.Report{
		ProductId: uint64(in.GetId()),
		UserId:    uint64(in.GetUserId()),
		Reason:    in.GetReason(),
		Description: sql.NullString{
			String: in.GetDescription(),
			Valid:  in.GetDescription() != "",
		},
		Images: reportImages,
	}

	result, err := l.svcCtx.ReportModel.Insert(l.ctx, &reportData)
	if err != nil {
		l.Errorf("Failed to insert report for product ID %d by user ID %d: %v", in.GetId(), in.GetUserId(), err)
		return nil, err
	}

	newReportId, err := result.LastInsertId()
	if err != nil {
		l.Errorf("Failed to get last insert ID for report: %v", err)
		return nil, err
	}

	return &product.ReportProductResponse{
		ReportId: newReportId,
		Message:  "举报成功",
	}, nil
}
