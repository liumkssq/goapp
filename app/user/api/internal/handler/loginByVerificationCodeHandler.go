package handler

import (
	"net/http"

	"github.com/liumkssq/goapp/app/user/api/internal/logic"
	"github.com/liumkssq/goapp/app/user/api/internal/svc"
	"github.com/liumkssq/goapp/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 验证码登录
func LoginByVerificationCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByVerificationCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginByVerificationCodeLogic(r.Context(), svcCtx)
		resp, err := l.LoginByVerificationCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
