package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLostFoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLostFoundLogic {
	return &CommentLostFoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论失物招领
func (l *CommentLostFoundLogic) CommentLostFound(in *lostfound.CommentLostFoundRequest) (*lostfound.CommentLostFoundResponse, error) {
	// todo: add your logic here and delete this line

	return &lostfound.CommentLostFoundResponse{}, nil
}
