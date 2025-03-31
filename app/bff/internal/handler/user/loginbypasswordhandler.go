package user

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/user"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 账号密码登录
func LoginByPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginByPasswordLogic(r.Context(), svcCtx)
		resp, err := l.LoginByPassword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
