package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/liumkssq/goapp/app/im/ws/internal/svc"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"strconv"
)

type JwtAuth struct {
	svc    *svc.ServiceContext
	parser *token.TokenParser
	logx.Logger
}

func (j *JwtAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	tok, err := j.parser.ParseToken(r, j.svc.Config.Auth.AccessSecret, "")
	if err != nil {
		j.Errorf("parse token err %v ", err)
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
	return strconv.FormatInt(ctxdata.GetUidFromCtx(r.Context()), 10)
}

func NewJwtAuth(svc *svc.ServiceContext) *JwtAuth {
	return &JwtAuth{
		svc:    svc,
		parser: token.NewTokenParser(),
		Logger: logx.WithContext(context.Background()),
	}
}
