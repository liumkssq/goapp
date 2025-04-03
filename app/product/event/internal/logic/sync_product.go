package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/product/event/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SyncProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncProductLogic(ctx context.Context, svcCtx *svc.ServiceContext, logger logx.Logger) *SyncProductLogic {
	return &SyncProductLogic{ctx: ctx, svcCtx: svcCtx, Logger: logger}
}

// SyncProduct 商品同步
func (l *SyncProductLogic) ProductProducer() error {

}
