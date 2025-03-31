package user

import (
	"net/http"

	"github.com/liumkssq/goapp/app/bff/internal/logic/user"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 发送验证码
func SendVerificationCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendVerificationCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSendVerificationCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendVerificationCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
