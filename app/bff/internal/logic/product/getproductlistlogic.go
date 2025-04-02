package product

import (
	"context"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListLogic) GetProductList(req *types.ProductListReq) (resp *types.ProductListResp, err error) {
	// todo: add your logic here and delete this line
	productList, err := l.svcCtx.ProductRpc.GetProductList(l.ctx, &product.GetProductListRequest{
		Page:       req.Page,
		Limit:      req.Limit,
		CategoryId: req.CategoryId,
		Keyword:    req.Keyword,
		Sort:       req.Sort,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	list := make([]types.Product, 0)
	for _, v := range productList.List {
		list = append(list, types.Product{
			Id:            v.Id,
			Title:         v.Title,
			Description:   v.Description,
			Price:         v.Price,
			OriginalPrice: v.OriginalPrice,
			ViewCount:     v.ViewCount,
			LikeCount:     v.LikeCount,
			CommentCount:  v.CommentCount,
			CategoryId:    v.CategoryId,
			CategoryName:  v.CategoryName,
			Condition:     v.Condition,
			ContactInfo:   v.ContactInfo,
			Status:        string(v.Status),
			ContactWay:    v.ContactWay,
			Location:      v.Location,
			Images:        v.Images,
			SellerId:      v.SellerId,
			SellerName:    v.SellerName,
			SellerAvatar:  v.SellerAvatar,
		})
	}
	resp = &types.ProductListResp{
		Page:  productList.Page,
		Limit: productList.Limit,
		Total: productList.Total,
		List:  list,
	}

	return resp, nil
}
