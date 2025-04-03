package allsearch

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLostFoundLogic {
	return &SearchLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLostFoundLogic) SearchLostFound(req *types.SearchLostFoundReq) (resp *types.SearchLostFoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
