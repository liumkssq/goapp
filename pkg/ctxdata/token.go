package ctxdata

import (
	"context"
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

const Identify = "liumkssq"

func GetJWTToken(secretKey string, iat, seconds int64, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64

	// 1. 直接从ctx中尝试获取
	if jsonUid, ok := ctx.Value(Identify).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			return int64Uid
		}
	}

	// 2. 从go-zero的JWT中间件设置的值中获取
	if claims, ok := ctx.Value("JWT_CLAIMS").(jwt.MapClaims); ok {
		if idStr, ok := claims[Identify].(float64); ok {
			return int64(idStr)
		}
	}

	return uid
}
