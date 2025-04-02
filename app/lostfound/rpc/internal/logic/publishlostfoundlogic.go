package logic

import (
	"context"
	"database/sql"
	"github.com/liumkssq/goapp/app/lostfound/model"
	"time"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLostFoundLogic {
	return &PublishLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布失物招领
func (l *PublishLostFoundLogic) PublishLostFound(
	in *lostfound.PublishLostFoundRequest) (*lostfound.PublishLostFoundResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 获取用户信息
	var LostType string
	if in.Type == lostfound.LostFoundType_LOSTT {
		LostType = "lost"
	} else {
		LostType = "found"
	}
	var images string
	for _, image := range in.Images {
		images += image + ","
	}
	var tags string
	for _, tag := range in.Tags {
		tags += tag + ","
	}
	//convert string to time
	//var lostFoundTime time.Time
	lostFoundTime, _ := time.Parse("2006-01-02 15:04:05", in.LostFoundTime)
	lostfoundItem := &model.LostFoundItem{
		Title:       in.Title,
		Description: in.Description,
		Type:        LostType,
		Category: sql.NullString{
			String: in.Category,
			Valid:  true,
		},
		Location: sql.NullString{
			String: in.Location,
			Valid:  true,
		},
		LocationDetail: sql.NullString{},
		ContactInfo: sql.NullString{
			String: in.ContactInfo,
			Valid:  true,
		},
		ContactWay: sql.NullString{
			String: in.ContactWay,
			Valid:  true,
		},
		Images: sql.NullString{
			String: images,
			Valid:  true,
		},
		Status: "published",
		Tags: sql.NullString{
			String: tags,
			Valid:  true,
		},
		RewardInfo: sql.NullString{
			String: in.RewardInfo,
			Valid:  true,
		},
		LostFoundTime: sql.NullTime{
			Time:  lostFoundTime,
			Valid: true,
		},
		//PublisherId: uint64(in.UserId),
		//PublisherAvatar: sql.NullString{
		//	String: in.im,
		//	Valid:  true,
		//},
	}

	result, err := l.svcCtx.LostFoundItemModel.Insert(l.ctx, nil, lostfoundItem)
	if err != nil {
		return nil, err
	}
	resultId, err := result.LastInsertId()
	return &lostfound.PublishLostFoundResponse{
		Id:      resultId,
		Message: "发布成功",
	}, nil
}
