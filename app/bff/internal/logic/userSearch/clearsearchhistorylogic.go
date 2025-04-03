package userSearch

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearSearchHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearSearchHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearSearchHistoryLogic {
	return &ClearSearchHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearSearchHistoryLogic) ClearSearchHistory() (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
