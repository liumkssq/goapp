package userSearch

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSearchHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSearchHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSearchHistoryLogic {
	return &GetSearchHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSearchHistoryLogic) GetSearchHistory() (resp *types.SearchHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
