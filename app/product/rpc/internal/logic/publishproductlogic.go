package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/liumkssq/goapp/app/product/model"
	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishProductLogic {
	return &PublishProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布商品
func (l *PublishProductLogic) PublishProduct(in *product.PublishProductRequest) (*product.PublishProductResponse, error) {

	l.Logger.Info("PublishProduct request: ", in)
	//var images string
	//for _, image := range in.Images {
	//	images += image + ","
	//}
	imagesJSON, err := json.Marshal(in.Images)
	if err != nil {
		l.Logger.Error("Marshal images error: ", err)
		return nil, err
	}
	//var tags string
	//for _, tag := range in.Tags {
	//	tags += tag + ","
	//}
	// 修改为JSON处理：
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
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.UserId))
	if err != nil {
		l.Logger.Error("PublishProduct error: ", err)
		return nil, err
	}
	userName := userItem.Username
	result, err := l.svcCtx.ProductModel.Insert(l.ctx, nil, &model.Product{
		Title:       in.Title,
		Description: in.Description,
		Price:       in.Price,
		OriginalPrice: sql.NullFloat64{
			Float64: in.OriginalPrice,
			Valid:   true,
		},
		CategoryId: uint64(in.CategoryId),
		Images:     string(imagesJSON),
		Condition:  in.Condition,
		ContactInfo: sql.NullString{
			String: in.ContactInfo,
			Valid:  true,
		},
		ContactWay: sql.NullString{
			String: in.ContactWay,
			Valid:  true,
		},
		Location: sql.NullString{
			String: in.Location,
			Valid:  true,
		},
		LocationDetail: sql.NullString{},
		Tags: sql.NullString{
			String: tagsJSON,
			Valid:  true,
		},
		Status:     "active",
		SellerId:   uint64(in.UserId),
		SellerName: userName,
		SellerAvatar: sql.NullString{
			String: userItem.AvatarUrl.String,
			Valid:  true,
		},
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	})
	if err != nil {
		l.Logger.Error("PublishProduct error: ", err)
		return nil, err
	}

	productId, err := result.LastInsertId()

	if err != nil {
		l.Logger.Error("PublishProduct error: ", err)
		return nil, err
	}
	return &product.PublishProductResponse{
		Id: int64(productId),
	}, nil
}
