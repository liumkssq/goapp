package lostfound

import (
	"context"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLostFoundListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLostFoundListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLostFoundListLogic {
	return &GetLostFoundListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLostFoundListLogic) GetLostFoundList(req *types.LostFoundListReq) (*types.LostFoundListResp, error) {
	// 调用RPC服务获取失物招领列表
	lostFoundList, err := l.svcCtx.LostFoundRpc.GetLostFoundList(l.ctx, &lostfound.GetLostFoundListRequest{
		Page:       req.Page,
		Limit:      req.Limit,
		CategoryId: req.CategoryId,
		Keywords:   req.Keywords,
		Type:       req.Type,
		Sort:       req.Sort,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	// 转换RPC响应到BFF响应
	list := make([]types.LostFoundItem, 0)
	for _, item := range lostFoundList.List {
		list = append(list, types.LostFoundItem{
			Id:              item.Id,
			Title:           item.Title,
			Description:     item.Description,
			Type:            string(item.Type),
			Category:        item.Category,
			Location:        item.Location,
			ContactInfo:     item.ContactInfo,
			ContactWay:      item.ContactWay,
			Images:          item.Images,
			Status:          string(item.Status),
			Tags:            item.Tags,
			RewardInfo:      item.RewardInfo,
			LostFoundTime:   item.LostFoundTime,
			PublisherName:   item.PublisherName,
			PublisherAvatar: item.PublisherAvatar,
			ViewCount:       item.ViewCount,
			LikeCount:       item.LikeCount,
			CommentCount:    item.CommentCount,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}

	// 构建响应
	var resp = &types.LostFoundListResp{
		List:  list,
		Page:  lostFoundList.Page,
		Limit: lostFoundList.Limit,
		Total: lostFoundList.Total,
	}

	return resp, nil
}
