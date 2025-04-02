package product

import (
	"context"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailLogic {
	return &GetProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductDetailLogic) GetProductDetail(req *types.ProductDetailReq) (*types.ProductDetailResp, error) {
	// todo: add your logic here and delete this line
	logx.Infof("GetProductDetail req: %v", req)
	//l.svcCtx.ProductRpc.GetProductDetail(l.ctx)
	productDetail, err := l.svcCtx.ProductRpc.GetProductDetail(l.ctx, &product.GetProductDetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var status string
	if productDetail.Product.Status == product.ProductStatus_ACTIVE {
		status = "ACTIVE"
	}
	resp := &types.ProductDetailResp{
		Product: types.Product{
			Id:            productDetail.Product.Id,
			Title:         productDetail.Product.Title,
			Description:   productDetail.Product.Description,
			Price:         productDetail.Product.Price,
			OriginalPrice: productDetail.Product.OriginalPrice,
			Images:        productDetail.Product.Images,
			Condition:     productDetail.Product.Condition,
			ContactInfo:   productDetail.Product.ContactInfo,
			ContactWay:    productDetail.Product.ContactWay,
			Location:      productDetail.Product.Location,
			//LocationDetail: productDetail.Product.LocationDetail,
			Tags:         productDetail.Product.Tags,
			Status:       status,
			SellerId:     productDetail.Product.SellerId,
			SellerName:   productDetail.Product.SellerName,
			SellerAvatar: productDetail.Product.SellerAvatar,
			ViewCount:    productDetail.Product.ViewCount,
			LikeCount:    productDetail.Product.LikeCount,
			CommentCount: productDetail.Product.CommentCount,
		},
	}
	return resp, nil
}
