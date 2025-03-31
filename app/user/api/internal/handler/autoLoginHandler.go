package handler

import (
	"net/http"

	"github.com/liumkssq/goapp/app/user/api/internal/logic"
	"github.com/liumkssq/goapp/app/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 自动登录（仅用于开发环境测试）
func AutoLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAutoLoginLogic(r.Context(), svcCtx)
		resp, err := l.AutoLogin()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
