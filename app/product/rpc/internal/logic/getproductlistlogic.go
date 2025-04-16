package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"strconv"

	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品列表
func (l *GetProductListLogic) GetProductList(in *product.GetProductListRequest) (*product.GetProductListResponse, error) {

	rowBuilder := squirrel.Select().From("`product`")

	// 添加过滤条件
	if in.CategoryId > 0 {
		rowBuilder = rowBuilder.Where("`category_id` = ?", in.CategoryId)
	}

	status, _ := strconv.Atoi(in.Status)
	if status > 0 {
		rowBuilder = rowBuilder.Where("`status` = ?", status)
	}
	if len(in.Keyword) > 0 {
		rowBuilder = rowBuilder.Where("`name` LIKE ?", "%"+in.Keyword+"%")
	}
	// 设置分页参数
	page := in.Page
	if page < 1 {
		page = 1
	}
	pageSize := in.Limit
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 设置排序
	orderBy := ""
	if in.Sort != "" {
		// 将前端排序参数映射到有效的数据库列
		switch in.Sort {
		case "hot":
			// 可以使用视图计数作为热门指标
			orderBy = "`view_count` DESC"
		case "new":
			orderBy = "`created_at` DESC"
		case "price_asc":
			orderBy = "`price` ASC"
		case "price_desc":
			orderBy = "`price` DESC"
		default:
			// 默认排序
			orderBy = "`id` DESC"
		}
	}

	// 查询产品列表
	products, err := l.svcCtx.ProductModel.FindPageListByPage(l.ctx, rowBuilder, page, pageSize, orderBy)
	if err != nil {
		return nil, err
	}
	productList := make([]*product.Product, 0, len(products))
	for _, p := range products {
		var oPrice float64
		if p.OriginalPrice.Valid {
			oPrice = p.OriginalPrice.Float64
		} else {
			oPrice = p.Price
		}
		productList = append(productList, &product.Product{
			Id:            int64(p.Id),
			Title:         p.Title,
			Description:   p.Description,
			Price:         p.Price,
			OriginalPrice: oPrice,
			CategoryId:    int64(p.CategoryId),
			Images:        []string{p.Images},
			ContactInfo:   p.ContactInfo.String,
			ContactWay:    p.ContactWay.String,
			Location:      p.Location.String,
			Tags:          []string{p.Tags.String},
			Status:        product.ProductStatus(int32(status)),
			SellerId:      int64(p.SellerId),
			SellerName:    p.SellerName,
			SellerAvatar:  p.SellerAvatar.String,
			ViewCount:     int64(p.ViewCount),
			LikeCount:     int64(p.LikeCount),
			CommentCount:  int64(p.CommentCount),
		})
	}

	return &product.GetProductListResponse{
		List:  productList,
		Total: int64(len(productList)),
		Page:  in.Page,
		Limit: in.Limit,
	}, nil
}
