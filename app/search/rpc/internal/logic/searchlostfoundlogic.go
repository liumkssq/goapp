package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/search/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLostFoundLogic {
	return &SearchLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索失物招领
func (l *SearchLostFoundLogic) SearchLostFound(in *search.SearchLostFoundRequest) (*search.SearchLostFoundResponse, error) {
	// todo: add your logic here and delete this line

	return &search.SearchLostFoundResponse{}, nil
}
