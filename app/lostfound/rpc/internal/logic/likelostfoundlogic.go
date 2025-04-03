package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"
	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLostFoundLogic {
	return &LikeLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞失物招领
func (l *LikeLostFoundLogic) LikeLostFound(in *lostfound.LikeLostFoundRequest) (*lostfound.LikeLostFoundResponse, error) {
	// todo: add your logic here and delete this line

	return &lostfound.LikeLostFoundResponse{}, nil
}
