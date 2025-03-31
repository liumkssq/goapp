package ctxdata

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetUidFromCtx(t *testing.T) {
	userId := int64(1)
	ctx := context.WithValue(context.Background(), CtxKeyJwtUserId, json.Number(fmt.Sprintf("%d", userId)))
	uid := GetUidFromCtx(ctx)
	fmt.Printf("uid : %d", uid)
}
