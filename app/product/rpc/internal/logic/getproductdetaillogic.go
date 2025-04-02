package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailLogic {
	return &GetProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品详情
func (l *GetProductDetailLogic) GetProductDetail(in *product.GetProductDetailRequest) (*product.GetProductDetailResponse, error) {
	// todo: add your logic here and delete this line

	productItem, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	logx.Infof("GetProductDetail productItem: %v", productItem)
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("GetProductDetail request: %v", in)
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, productItem.SellerId)
	logx.Infof("GetProductDetail userItem: %v", userItem)
	if err != nil {
		return nil, err
	}
	var resp product.GetProductDetailResponse
	//status := productItem.Status
	resp.Product = &product.Product{
		Id:            int64(productItem.Id),
		Title:         productItem.Title,
		Price:         productItem.Price,
		OriginalPrice: productItem.OriginalPrice.Float64,
		Description:   productItem.Description,
		Images:        []string{productItem.Images},
		SellerId:      int64(productItem.SellerId),
		SellerName:    userItem.Username,
		SellerAvatar:  userItem.AvatarUrl.String,
		ViewCount:     int64(productItem.ViewCount),
		LikeCount:     int64(productItem.LikeCount),
		CommentCount:  int64(productItem.CommentCount),
	}
	return &resp, nil
}
