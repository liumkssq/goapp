package user

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/user"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 自动登录（仅用于开发环境测试）
func AutoLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewAutoLoginLogic(r.Context(), svcCtx)
		resp, err := l.AutoLogin()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
