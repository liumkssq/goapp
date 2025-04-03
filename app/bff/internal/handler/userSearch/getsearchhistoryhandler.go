package userSearch

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/userSearch"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSearchHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userSearch.NewGetSearchHistoryLogic(r.Context(), svcCtx)
		resp, err := l.GetSearchHistory()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
