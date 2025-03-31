package handler

import (
	"net/http"

	"github.com/liumkssq/goapp/app/user/api/internal/logic"
	"github.com/liumkssq/goapp/app/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 退出登录
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
