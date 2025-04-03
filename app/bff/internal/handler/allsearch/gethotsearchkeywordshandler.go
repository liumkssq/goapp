package allsearch

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/allsearch"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHotSearchKeywordsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := allsearch.NewGetHotSearchKeywordsLogic(r.Context(), svcCtx)
		resp, err := l.GetHotSearchKeywords()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
