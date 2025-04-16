package im

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/im"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 建立会话
func SetUpUserConversationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUpUserConversationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := im.NewSetUpUserConversationLogic(r.Context(), svcCtx)
		resp, err := l.SetUpUserConversation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
