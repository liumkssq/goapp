package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/search/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/search/rpc/search"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
)

type GlobalSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGlobalSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GlobalSearchLogic {
	return &GlobalSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 全局搜索
func (l *GlobalSearchLogic) GlobalSearch(in *search.GlobalSearchRequest) (*search.GlobalSearchResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 解析参数 split " " 空格 todo 其他后续优化先简单用空格
	searchType := in.Type
	var resp search.GlobalSearchResponse
	keywords := strings.Split(in.Keyword, " ")
	if len(keywords) != 0 {
		in.Keyword = keywords[0]
	} else {
		in.Keyword = ""
	}
	if searchType == "all" || searchType == "product" {
		searchProduct, err := l.svcCtx.ProductModel.SearchProduct(l.ctx, in.Keyword, searchType, in.Page, in.Limit)
		if err != nil {
			return nil, err
		}
		for _, item := range searchProduct {
			ProductSearchResp := &search.SearchResult{
				Type:    searchType,
				Id:      int64(item.Id),
				Title:   item.Title,
				Content: item.Description,
				Image:   item.Images,
				Extra: map[string]string{
					"price": strconv.Itoa(int(item.Price)),
				},
			}
			resp.List = append(resp.List, ProductSearchResp)
		}
	}
	if searchType == "all" || searchType == "lostfound" {
		// 2. 调用 lostfound rpc
		// 3. 调用 user rpc
		searchLostFound, err := l.svcCtx.LostFoundModel.SearchLostFound(l.ctx, in.Keyword, searchType, in.Page, in.Limit)
		if err != nil {
			return nil, err
		}
		for _, item := range searchLostFound {
			LostFoundSearchResp := &search.SearchResult{
				Type:    searchType,
				Id:      int64(item.Id),
				Title:   item.Title,
				Content: item.Description,
				Image:   item.Images.String,
				Extra: map[string]string{
					"status": item.Status,
				},
			}
			resp.List = append(resp.List, LostFoundSearchResp)
			//
			//var eg errgroup.Group
			//eg.Go(func() error {
			//	l.svcCtx.
			//})

		}
	}
	if searchType == "all" || searchType == "user" {

		searchUser, err := l.svcCtx.UserModel.SearchUser(l.ctx, in.Keyword, searchType, in.Page, in.Limit)
		if err != nil {
			return nil, err
		}
		for _, item := range searchUser {
			UserSearchResp := &search.SearchResult{
				Type:    searchType,
				Id:      int64(item.Id),
				Title:   item.Username,
				Content: item.Nickname.String,
				Image:   item.AvatarUrl.String,
				Extra: map[string]string{
					"phone": item.Phone,
					"bio":   item.Bio.String,
				},
			}
			resp.List = append(resp.List, UserSearchResp)
		}
	}
	resp.Total = int64(len(resp.List))
	return &resp, nil
}
