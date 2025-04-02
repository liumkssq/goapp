package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/search/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/search/rpc/search"
	"golang.org/x/sync/errgroup"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
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
	keywords := strings.Split(in.Keyword, " ")
	var searchType string
	if len(keywords) > 0 {
		searchType = keywords[0]
	} else {
		searchType = "all"
	}
	var eg errgroup.Group
	eg.Go(func() error {
		// 2. 调用对应的 repo
		//switch searchType {
		//case "user":
		//	// 3. 调用 user repo
		//	users, err := l.svcCtx.UserRepo.SearchUser(l.ctx, keywords)
		//	if err != nil {
		//		return err
		//	}
		//	l.Logger.Infof("users: %v", users)
		//case "article":
		//	// 4. 调用 article repo
		//	articles, err := l.svcCtx.ArticleRepo.SearchArticle(l.ctx, in.Uid, keywords)
		//	if err != nil {
		//		return err
		//	}
		//	l.Logger.Infof("articles: %v", articles)
		//default:
		//	l.Logger.Infof("searchType: %s", searchType)
		//}
		//return nil
		return nil
	})

	return &search.GlobalSearchResponse{}, nil
}

//func (l *GlobalSearchLogic) parseExpr(expr string) (string, []domain.QueryMeta, error) {
//	searchParams := strings.SplitN(expr, ":", 3)
//	if len(searchParams) == 3 {
//		typ := searchParams[0]
//		if typ != "biz" {
//			return "", nil, errors.New("参数错误")
//		}
//		biz := searchParams[1]
//		keywords := searchParams[2]
//		return biz, s.getQueryMeta(keywords), nil
//	}
//	return "", nil, errors.New("参数错误")
//}
