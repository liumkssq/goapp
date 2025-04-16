package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/liumkssq/goapp/app/im/ws/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"reflect"
	"testing"
)

func TestToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4NjcyODcsImlhdCI6MTc0NDc4MDg4NywiaW1vb2MuY29tIjoiMSJ9.j9EO_FhKtevLdi5vnXKI45Kws6FePiCV4nzJzBSeI2U"
	secretKey := "imooc.com"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return
	}

	if token.Valid {
		fmt.Println("令牌有效")
	}
}

func TestJwtAuth_UserId(t *testing.T) {
	type fields struct {
		svc    *svc.ServiceContext
		parser *token.TokenParser
		Logger logx.Logger
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JwtAuth{
				svc:    tt.fields.svc,
				parser: tt.fields.parser,
				Logger: tt.fields.Logger,
			}
			if got := j.UserId(tt.args.r); got != tt.want {
				t.Errorf("UserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJwtAuth(t *testing.T) {
	type args struct {
		svc *svc.ServiceContext
	}
	tests := []struct {
		name string
		args args
		want *JwtAuth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJwtAuth(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJwtAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
