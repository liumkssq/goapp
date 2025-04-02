package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/search/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchArticlesLogic {
	return &SearchArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索文章
func (l *SearchArticlesLogic) SearchArticles(in *search.SearchArticlesRequest) (*search.SearchArticlesResponse, error) {
	// todo: add your logic here and delete this line

	return &search.SearchArticlesResponse{}, nil
}
