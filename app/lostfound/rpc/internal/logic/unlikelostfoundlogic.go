package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnlikeLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeLostFoundLogic {
	return &UnlikeLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消点赞失物招领
func (l *UnlikeLostFoundLogic) UnlikeLostFound(in *lostfound.UnlikeLostFoundRequest) (*lostfound.UnlikeLostFoundResponse, error) {
	// todo: add your logic here and delete this line

	return &lostfound.UnlikeLostFoundResponse{}, nil
}
