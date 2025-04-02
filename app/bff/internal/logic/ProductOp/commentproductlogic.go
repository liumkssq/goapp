package ProductOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentProductLogic {
	return &CommentProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentProductLogic) CommentProduct(req *types.CommentProductReq) (resp *types.CommentProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
