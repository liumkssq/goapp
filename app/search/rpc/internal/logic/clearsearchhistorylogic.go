package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/search/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearSearchHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearSearchHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearSearchHistoryLogic {
	return &ClearSearchHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清除搜索历史
func (l *ClearSearchHistoryLogic) ClearSearchHistory(in *search.ClearSearchHistoryRequest) (*search.ClearSearchHistoryResponse, error) {
	// todo: add your logic here and delete this line

	return &search.ClearSearchHistoryResponse{}, nil
}
