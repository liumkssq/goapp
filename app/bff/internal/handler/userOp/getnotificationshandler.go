package userOp

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/userOp"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取通知
func GetNotificationsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userOp.NewGetNotificationsLogic(r.Context(), svcCtx)
		resp, err := l.GetNotifications()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
