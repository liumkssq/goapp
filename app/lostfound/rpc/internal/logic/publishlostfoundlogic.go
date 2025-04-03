package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
	imagesJSON, err := json.Marshal(in.Images)
	if err != nil {
		l.Logger.Error("Marshal images error: ", err)
		return nil, err
	}
	var tagsJSON string
	if len(in.Tags) > 0 {
		tagsData, err := json.Marshal(in.Tags)
		if err != nil {
			l.Logger.Error("Marshal tags error: ", err)
			return nil, err
		}
		tagsJSON = string(tagsData)
	} else {
		tagsJSON = "[]" // 空数组的JSON表示
	}
	//convert string to time
	//var lostFoundTime time.Time

	fmt.Printf("in: %v", in.UserId)
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.UserId))
	if err != nil {
		l.Logger.Error("PublishProduct error: ", err)
		return nil, err
	}
	logx.Infof("userItem: %v", userItem)
	userName := userItem.Username
	var lostFoundTime time.Time
	var valid bool

	// 有时间值且格式正确时才设置为有效
	if in.LostFoundTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", in.LostFoundTime)
		if err == nil {
			lostFoundTime = parsedTime
			valid = true
		} else {
			// 记录错误但不中断
			l.Logger.Error("Parse lost_found_time error:", err)
		}
	}

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
			String: string(imagesJSON),
			Valid:  true,
		},
		Status: "pending",
		Tags: sql.NullString{
			String: tagsJSON,
			Valid:  true,
		},
		RewardInfo: sql.NullString{
			String: in.RewardInfo,
			Valid:  true,
		},
		LostFoundTime: sql.NullTime{
			Time:  lostFoundTime,
			Valid: valid,
		},
		PublisherId:   userItem.Id,
		PublisherName: userName,
		PublisherAvatar: sql.NullString{
			String: userItem.AvatarUrl.String,
			Valid:  true,
		},
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
	}

	result, err := l.svcCtx.LostFoundItemModel.Insert(l.ctx, nil, lostfoundItem)
	if err != nil {
		logx.Error("PublishLostFound error: ", err)
		return nil, err
	}
	resultId, err := result.LastInsertId()
	return &lostfound.PublishLostFoundResponse{
		Id:      resultId,
		Message: "发布成功",
	}, nil
}
