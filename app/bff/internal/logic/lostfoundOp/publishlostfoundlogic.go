package lostfoundOp

import (
	"context"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLostFoundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLostFoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLostFoundLogic {
	return &PublishLostFoundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLostFoundLogic) PublishLostFound(req *types.PublishLostFoundReq) (*types.PublishLostFoundResp, error) {
	var PType lostfound.LostFoundType
	if req.Type == "lost" {
		PType = lostfound.LostFoundType_LOSTT
	} else {
		PType = lostfound.LostFoundType_FOUNDT
	}

	lostFound, err := l.svcCtx.LostFoundRpc.PublishLostFound(l.ctx, &lostfound.PublishLostFoundRequest{
		Title:         req.Title,
		Description:   req.Description,
		Location:      req.Location,
		ContactInfo:   req.ContactInfo,
		ContactWay:    req.ContactWay,
		Images:        req.Images,
		Type:          PType,
		Category:      req.Category,
		Tags:          req.Tags,
		LostFoundTime: req.LostFoundTime,
		RewardInfo:    req.RewardInfo,
	})
	if err != nil {
		logx.Error("PublishLostFound failed", err)
		return nil, err
	}
	resp := types.PublishLostFoundResp{
		Id:      lostFound.Id,
		Message: "发布成功",
	}
	return &resp, nil
}
