package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLostFoundCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLostFoundCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLostFoundCommentsLogic {
	return &GetLostFoundCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取失物招领评论
func (l *GetLostFoundCommentsLogic) GetLostFoundComments(in *lostfound.GetLostFoundCommentsRequest) (*lostfound.GetLostFoundCommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &lostfound.GetLostFoundCommentsResponse{}, nil
}
