/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/liumkssq/goapp/app/im/ws/internal/svc"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
)

type JwtAuth struct {
	svc    *svc.ServiceContext
	parser *token.TokenParser
	logx.Logger
}

func NewJwtAuth(svc *svc.ServiceContext) *JwtAuth {
	return &JwtAuth{
		svc:    svc,
		parser: token.NewTokenParser(),
		Logger: logx.WithContext(context.Background()),
	}
}

func (j *JwtAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("sec-websocket-protocol")
	j.Infof("接收到的token: %s\n", token)
	j.Infof("使用的密钥: %s\n", j.svc.Config.JwtAuth.AccessSecret)

	if token != "" {
		fmt.Println("token", token)
		r.Header.Set("Authorization", token)
	}
	tok, err := j.parser.ParseToken(r, j.svc.Config.JwtAuth.AccessSecret, "")
	if err != nil {
		j.Errorf("解析token错误: %v", err)
		return false
	}

	if !tok.Valid {
		return false
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	*r = *r.WithContext(context.WithValue(r.Context(), ctxdata.Identify, claims[ctxdata.Identify]))

	return true
}

func (j *JwtAuth) UserId(r *http.Request) string {
	return ctxdata.GetUId(r.Context())
}
