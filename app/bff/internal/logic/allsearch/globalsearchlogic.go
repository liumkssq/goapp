package allsearch

import (
	"context"
	"fmt"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/liumkssq/goapp/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type GlobalSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGlobalSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GlobalSearchLogic {
	return &GlobalSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GlobalSearchLogic) GlobalSearch(req *types.GlobalSearchReq) (*types.GlobalSearchResp, error) {
	// todo: add your logic here and delete this line
	logx.Infof("GlobalSearch req: %v", req)
	globalSearchResponse, err := l.svcCtx.SearchRpc.GlobalSearch(l.ctx, &search.GlobalSearchRequest{
		Keyword: req.Keyword,
		Type:    req.Type,
		Page:    req.Page,
		Limit:   req.Limit,
	})
	if err != nil {
		return nil, err
	}
	// 假设 globalSearchResponse.List 是 []*SearchResult 类型
	var resultList []types.SearchResult

	for _, item := range globalSearchResponse.List {
		if item != nil { // 确保指针非空
			resultList = append(resultList, types.SearchResult{
				Type:      item.Type,
				Id:        item.Id,
				Title:     item.Title,
				Content:   item.Content,
				Image:     item.Image,
				CreatedAt: item.CreatedAt,
			})
		}
	}
	fmt.Printf("resultList: %v", resultList)
	logx.Infof("GlobalSearchResp: %v", globalSearchResponse)
	var resp *types.GlobalSearchResp
	resp = &types.GlobalSearchResp{
		List:  resultList,
		Total: globalSearchResponse.Total,
		Page:  req.Page,
		Limit: req.Limit,
	}
	logx.Infof("GlobalSearchResp: %v", resp)

	//globalSearchResponse.List   []*SearchResult
	//types.GlobalSearchResp{
	//	List: make([]types.SearchResult, 0), //
	//}
	//*
	//
	return resp, nil
}
