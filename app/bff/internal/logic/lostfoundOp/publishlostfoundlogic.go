package lostfoundOp

import (
	"context"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"
	"github.com/liumkssq/goapp/pkg/ctxdata"

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
	uid := ctxdata.GetUidFromCtx(l.ctx)
	var PType lostfound.LostFoundType
	if req.Type == "lost" {
		PType = lostfound.LostFoundType_LOSTT
	} else {
		PType = lostfound.LostFoundType_FOUNDT
	}

	lostFound, err := l.svcCtx.LostFoundRpc.PublishLostFound(l.ctx, &lostfound.PublishLostFoundRequest{
		UserId:      uid,
		Title:       req.Title,
		Description: req.Description,
		Type:        PType,
		Category:    req.Category,
		Location:    req.Location,
		LocationDetail: map[string]string{
			"test": "test",
		},
		ContactInfo:   req.ContactInfo,
		ContactWay:    req.ContactWay,
		Images:        req.Images,
		Tags:          req.Tags,
		RewardInfo:    req.RewardInfo,
		LostFoundTime: req.LostFoundTime,
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
