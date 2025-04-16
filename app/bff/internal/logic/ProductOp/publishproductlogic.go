package ProductOp

import (
	"context"
	"fmt"
	"strconv"

	"github.com/liumkssq/goapp/app/product/rpc/product"
	"github.com/liumkssq/goapp/pkg/ctxdata"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishProductLogic {
	return &PublishProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishProductLogic) PublishProduct(req *types.PublishProductReq) (*types.PublishProductResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("PublishProduct")
	l.Logger.Infof("PublishProduct req: %v", req)
	userId := ctxdata.GetUId(l.ctx)
	userIdI, _ := strconv.Atoi(userId)
	fmt.Printf("condition: %v", req.Condition)
	fmt.Printf("PublishProduct userId: %v", userId)
	codition := mapCondition(req.Condition)
	publishProductResponse, err := l.svcCtx.ProductRpc.PublishProduct(l.ctx, &product.PublishProductRequest{
		UserId:         int64(userIdI),
		Title:          req.Title,
		Description:    req.Description,
		Price:          req.Price,
		OriginalPrice:  req.OriginalPrice,
		CategoryId:     req.CategoryId,
		Images:         req.Images,
		Condition:      codition,
		ContactInfo:    req.ContactInfo,
		ContactWay:     req.ContactWay,
		Location:       req.Location,
		LocationDetail: nil,
		Tags:           req.Tags,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.PublishProductResp{
		Id:      publishProductResponse.Id,
		Message: "发布成功",
	}
	return resp, nil
}

func mapCondition(frontendCondition string) string {
	switch frontendCondition {
	case "newProduct":
		return "全新"
	case "like_newProduct":
		return "9成新"
	case "goodProduct":
		return "8成新"
	case "slight_usedProduct":
		return "7成新"
	default:
		return "其他"
	}
}
