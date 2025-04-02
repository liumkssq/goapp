package lostfoundOp

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLostFoundLogic {
	return &CommentLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLostFoundLogic) CommentLostFound(req *types.CommentLostFoundReq) (resp *types.CommentLostFoundResp, err error) {
	// todo: add your logic here and delete this line

	return
}
