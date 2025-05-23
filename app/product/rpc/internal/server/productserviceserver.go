// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: product.proto

package server

import (
	"context"

	"github.com/liumkssq/goapp/app/product/rpc/internal/logic"
	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"
)

type ProductServiceServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServiceServer
}

func NewProductServiceServer(svcCtx *svc.ServiceContext) *ProductServiceServer {
	return &ProductServiceServer{
		svcCtx: svcCtx,
	}
}

// 获取商品列表
func (s *ProductServiceServer) GetProductList(ctx context.Context, in *product.GetProductListRequest) (*product.GetProductListResponse, error) {
	l := logic.NewGetProductListLogic(ctx, s.svcCtx)
	return l.GetProductList(in)
}

// 获取商品详情
func (s *ProductServiceServer) GetProductDetail(ctx context.Context, in *product.GetProductDetailRequest) (*product.GetProductDetailResponse, error) {
	l := logic.NewGetProductDetailLogic(ctx, s.svcCtx)
	return l.GetProductDetail(in)
}

// 获取商品分类
func (s *ProductServiceServer) GetProductCategories(ctx context.Context, in *product.GetProductCategoriesRequest) (*product.GetProductCategoriesResponse, error) {
	l := logic.NewGetProductCategoriesLogic(ctx, s.svcCtx)
	return l.GetProductCategories(in)
}

// 发布商品
func (s *ProductServiceServer) PublishProduct(ctx context.Context, in *product.PublishProductRequest) (*product.PublishProductResponse, error) {
	l := logic.NewPublishProductLogic(ctx, s.svcCtx)
	return l.PublishProduct(in)
}

// 更新商品
func (s *ProductServiceServer) UpdateProduct(ctx context.Context, in *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	l := logic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

// 删除商品
func (s *ProductServiceServer) DeleteProduct(ctx context.Context, in *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	l := logic.NewDeleteProductLogic(ctx, s.svcCtx)
	return l.DeleteProduct(in)
}

// 获取用户的商品
func (s *ProductServiceServer) GetUserProducts(ctx context.Context, in *product.GetUserProductsRequest) (*product.GetUserProductsResponse, error) {
	l := logic.NewGetUserProductsLogic(ctx, s.svcCtx)
	return l.GetUserProducts(in)
}

// 收藏/取消收藏商品
func (s *ProductServiceServer) FavoriteProduct(ctx context.Context, in *product.FavoriteProductRequest) (*product.FavoriteProductResponse, error) {
	l := logic.NewFavoriteProductLogic(ctx, s.svcCtx)
	return l.FavoriteProduct(in)
}

// 获取收藏的商品
func (s *ProductServiceServer) GetFavoriteProducts(ctx context.Context, in *product.GetFavoriteProductsRequest) (*product.GetFavoriteProductsResponse, error) {
	l := logic.NewGetFavoriteProductsLogic(ctx, s.svcCtx)
	return l.GetFavoriteProducts(in)
}

// 举报商品
func (s *ProductServiceServer) ReportProduct(ctx context.Context, in *product.ReportProductRequest) (*product.ReportProductResponse, error) {
	l := logic.NewReportProductLogic(ctx, s.svcCtx)
	return l.ReportProduct(in)
}

// 评论商品
func (s *ProductServiceServer) CommentProduct(ctx context.Context, in *product.CommentProductRequest) (*product.CommentProductResponse, error) {
	l := logic.NewCommentProductLogic(ctx, s.svcCtx)
	return l.CommentProduct(in)
}

// 获取商品评论
func (s *ProductServiceServer) GetProductComments(ctx context.Context, in *product.GetProductCommentsRequest) (*product.GetProductCommentsResponse, error) {
	l := logic.NewGetProductCommentsLogic(ctx, s.svcCtx)
	return l.GetProductComments(in)
}
