package logic

import (
	"context"
	"fmt"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLostFoundDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLostFoundDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLostFoundDetailLogic {
	return &GetLostFoundDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取失物招领详情
func (l *GetLostFoundDetailLogic) GetLostFoundDetail(in *lostfound.GetLostFoundDetailRequest) (*lostfound.GetLostFoundDetailResponse, error) {
	lostFoundItem, err := l.svcCtx.LostFoundItemModel.FindOne(l.ctx, in.Id)
	logx.Infof("lostFoundItem: %v", lostFoundItem)
	if err != nil {
		return nil, err
	}
	fmt.Printf("lostFoundItem: %v", lostFoundItem.PublisherId)
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, lostFoundItem.PublisherId)
	if err != nil {
		logx.Errorf("GetLostFoundDetail FindOne userItem err: %v", err)
		return nil, err
	}
	var lostfoundType lostfound.LostFoundType
	if lostFoundItem.Type == "lost" {
		lostfoundType = lostfound.LostFoundType_LOSTT
	} else {
		lostfoundType = lostfound.LostFoundType_FOUNDT
	}
	var status lostfound.LostFoundStatus
	//状态：待处理、已找到、已认领、已关闭
	if lostFoundItem.Status == "待处理" {
		status = lostfound.LostFoundStatus_PENDING
	} else if lostFoundItem.Status == "已找到" {
		status = lostfound.LostFoundStatus_FOUND
	} else if lostFoundItem.Status == "已认领" {
		status = lostfound.LostFoundStatus_CLAIMED
	} else if lostFoundItem.Status == "已关闭" {
		status = lostfound.LostFoundStatus_CLOSED
	}
	lostFoundTime := lostFoundItem.LostFoundTime.Time.Format("2006-01-02 15:04:05")
	resp := &lostfound.GetLostFoundDetailResponse{
		Item: &lostfound.LostFoundItem{

			Id:          int64(lostFoundItem.Id),
			Title:       lostFoundItem.Title,
			Description: lostFoundItem.Description,
			Category:    lostFoundItem.Category.String,
			Location:    lostFoundItem.Location.String,
			//LocationDetail:  lostFoundItem.LocationDetail.String,
			ContactInfo: lostFoundItem.ContactInfo.String,
			Type:        lostfoundType,
			ContactWay:  lostFoundItem.ContactWay.String,
			Images:      []string{lostFoundItem.Images.String},
			Status:      status,
			Tags:        []string{lostFoundItem.Tags.String},
			RewardInfo:  lostFoundItem.RewardInfo.String,
			//LostFoundTime:   lostFoundItem.LostFoundTime.String,
			LostFoundTime:   lostFoundTime,
			PublisherId:     int64(lostFoundItem.PublisherId),
			PublisherName:   userItem.Username,
			PublisherAvatar: userItem.AvatarUrl.String,
			ViewCount:       int64(lostFoundItem.ViewCount),
			LikeCount:       int64(lostFoundItem.LikeCount),
			CommentCount:    int64(lostFoundItem.CommentCount),
		},
		Comments: nil,
	}
	logx.Infof("GetLostFoundDetail resp: %v", resp)

	return resp, nil
}
