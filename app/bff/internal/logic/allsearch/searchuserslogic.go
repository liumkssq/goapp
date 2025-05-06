package allsearch

import (
	"context"
	"github.com/liumkssq/goapp/app/search/rpc/searchservice"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUsersLogic {
	return &SearchUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUsersLogic) SearchUsers(req *types.SearchUsersReq) (*types.SearchUsersResp, error) {
	// todo: add your logic here and delete this line
	searchUsers, err := l.svcCtx.SearchRpc.SearchUsers(l.ctx, &searchservice.SearchUsersRequest{
		Keyword: req.Keyword,
		Page:    req.Page,
		Limit:   req.Limit,
	})
	var resp types.SearchUsersResp
	if err != nil {
		return nil, err
	}
	if searchUsers == nil {
		return nil, nil
	}
	resp.Total = searchUsers.Total
	for _, item := range searchUsers.List {
		userSearchResp := types.UserSearchResult{
			Id:       item.Id,
			Username: item.Username,
			Avatar:   item.Avatar,
		}
		resp.List = append(resp.List, userSearchResp)
	}

	return &resp, nil
}
