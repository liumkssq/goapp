package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLostFoundLogic {
	return &GetUserLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户发布的失物招领
func (l *GetUserLostFoundLogic) GetUserLostFound(in *lostfound.GetUserLostFoundRequest) (*lostfound.GetUserLostFoundResponse, error) {
	//l.svcCtx.UserModel.FindOne()
	//l.svcCtx.LostFoundItemModel.FindPageListByPage()

	return &lostfound.GetUserLostFoundResponse{}, nil
}
