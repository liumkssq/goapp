package userOp

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/userOp"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 标记通知为已读
func MarkNotificationAsReadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MarkNotificationAsReadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userOp.NewMarkNotificationAsReadLogic(r.Context(), svcCtx)
		resp, err := l.MarkNotificationAsRead(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
