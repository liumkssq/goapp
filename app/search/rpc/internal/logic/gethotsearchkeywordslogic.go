package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/search/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotSearchKeywordsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHotSearchKeywordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotSearchKeywordsLogic {
	return &GetHotSearchKeywordsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取热门搜索关键词
func (l *GetHotSearchKeywordsLogic) GetHotSearchKeywords(in *search.GetHotSearchKeywordsRequest) (*search.GetHotSearchKeywordsResponse, error) {
	// todo: add your logic here and delete this line

	return &search.GetHotSearchKeywordsResponse{}, nil
}
