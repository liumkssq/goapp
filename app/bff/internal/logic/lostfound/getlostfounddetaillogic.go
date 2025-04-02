package lostfound

import (
	"context"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLostFoundDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLostFoundDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLostFoundDetailLogic {
	return &GetLostFoundDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLostFoundDetailLogic) GetLostFoundDetail(req *types.LostFoundDetailReq) (resp *types.LostFoundDetailResp, err error) {
	// todo: add your logic here and delete this line
	//l.svcCtx.ProductRpc.GetProductDetail(l.ctx)
	lostFoundDetail, err := l.svcCtx.LostFoundRpc.GetLostFoundDetail(l.ctx, &lostfound.GetLostFoundDetailRequest{
		Id: req.Id,
	})
	logx.Infof("GetLostFoundDetail req: %v", req)
	if err != nil {
		return nil, err
	}
	resp = &types.LostFoundDetailResp{
		Item: types.LostFoundItem{
			Id:          lostFoundDetail.Item.Id,
			Title:       lostFoundDetail.Item.Title,
			Description: lostFoundDetail.Item.Description,
			//Type:            lostFoundDetail.Item.Type,
			//Status:          lostFoundDetail.Item.Status,
			ContactInfo:     lostFoundDetail.Item.ContactInfo,
			ContactWay:      lostFoundDetail.Item.ContactWay,
			PublisherId:     lostFoundDetail.Item.PublisherId,
			PublisherName:   lostFoundDetail.Item.PublisherName,
			PublisherAvatar: lostFoundDetail.Item.PublisherAvatar,
			CreatedAt:       lostFoundDetail.Item.CreatedAt,
			UpdatedAt:       lostFoundDetail.Item.UpdatedAt,
			ViewCount:       lostFoundDetail.Item.ViewCount,
		},
		Comments: nil,
	}
	return resp, nil
}
