package handler

import (
	"net/http"

	"github.com/liumkssq/goapp/app/user/api/internal/logic"
	"github.com/liumkssq/goapp/app/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取通知
func GetNotificationsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetNotificationsLogic(r.Context(), svcCtx)
		resp, err := l.GetNotifications()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
