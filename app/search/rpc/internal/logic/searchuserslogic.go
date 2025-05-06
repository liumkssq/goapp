package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/search/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUsersLogic {
	return &SearchUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索用户
func (l *SearchUsersLogic) SearchUsers(in *search.SearchUsersRequest) (*search.SearchUsersResponse, error) {
	//// todo: add your logic here and delete this line
	//// 1. 解析参数 split " " 空格 todo 其他后续优化先简单用空格
	//keywords := strings.Split(in.Keyword, " ")
	//if len(keywords) != 0 {
	//	in.Keyword = keywords[0]
	//} else {
	//	in.Keyword = ""
	//}
	//findUserResp, err := l.svcCtx.UserRpc.FindUser(l.ctx, &userservice.FindUserReq{
	//	//Name:  "",
	//	//Phone: "",
	//	Ids: []string{in.Keyword},
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if findUserResp == nil {
	//	return nil, nil
	//}
	//var resp search.SearchUsersResponse
	//resp.Total = int64(len(findUserResp.User))
	//for _, item := range findUserResp.User {
	//	UserSearchResp := &search.UserSearchResult{
	//		Id:       item.UserId,
	//		Username: item.Username,
	//		Avatar:   item.Avatar,
	//	}
	//	resp.List = append(resp.List, UserSearchResp)
	//}
	//return &resp, nil
	return nil, nil
}
