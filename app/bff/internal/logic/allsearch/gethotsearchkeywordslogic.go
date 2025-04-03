package allsearch

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotSearchKeywordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotSearchKeywordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotSearchKeywordsLogic {
	return &GetHotSearchKeywordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotSearchKeywordsLogic) GetHotSearchKeywords() (resp *types.HotSearchKeywordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
