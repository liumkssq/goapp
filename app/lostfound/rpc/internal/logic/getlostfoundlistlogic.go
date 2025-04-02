package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/liumkssq/goapp/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfound"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLostFoundListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLostFoundListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLostFoundListLogic {
	return &GetLostFoundListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取失物招领列表
func (l *GetLostFoundListLogic) GetLostFoundList(in *lostfound.GetLostFoundListRequest) (*lostfound.GetLostFoundListResponse, error) {
	rowBuilder := squirrel.Select().From("`lost_found_item`")
	// 添加过滤条件
	if in.CategoryId > 0 {
		rowBuilder = rowBuilder.Where("`category_id` = ?", in.CategoryId)
	}

	status, _ := strconv.Atoi(in.Status)
	if status > 0 {
		rowBuilder = rowBuilder.Where("`status` = ?", status)
	}

	if len(in.Keywords) > 0 {
		rowBuilder = rowBuilder.Where("`title` LIKE ?", "%"+in.Keywords+"%")
	}

	if len(in.Type) > 0 && in.Type != "all" {
		rowBuilder = rowBuilder.Where("`type` = ?", in.Type) // 失物招领类型：lost-丢失物品，found-拾得物品
	}

	// 设置分页参数
	page := in.Page
	if page < 1 {
		page = 1
	}

	pageSize := in.Limit
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 设置排序
	orderBy := ""
	if in.Sort != "" {
		// 将前端排序参数映射到有效的数据库列
		switch in.Sort {
		case "latest":
			orderBy = "`created_at` DESC"
		case "hot":
			orderBy = "`view_count` DESC"
		case "location":
			orderBy = "`location` ASC"
		default:
			// 默认排序
			orderBy = "`id` DESC"
		}
	}

	// 查询失物招领列表
	lostfounds, err := l.svcCtx.LostFoundItemModel.FindPageListByPage(l.ctx, rowBuilder, page, pageSize, orderBy)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "获取失物招领列表失败: %v", err)
	}

	fmt.Printf("lostfounds: %v", lostfounds)
	lostFoundList := make([]*lostfound.LostFoundItem, 0, len(lostfounds))
	for _, item := range lostfounds {
		// 处理图片和标签字符串，转换为数组
		var images []string
		if item.Images.Valid {
			images = append(images, strings.Split(item.Images.String, ",")...)
		}
		// 直接格式化time.Time类型
		createTime := item.CreatedAt.Format("2006-01-02 15:04:05")
		lostFoundTime := item.LostFoundTime.Time.Format("2006-01-02 15:04:05")
		publisherName := item.PublisherName
		publisherAvatar := item.PublisherAvatar.String

		lostFoundList = append(lostFoundList, &lostfound.LostFoundItem{
			Id:              int64(item.Id),
			Title:           item.Title,
			Description:     item.Description,
			Images:          images,
			ContactInfo:     item.ContactInfo.String,
			ContactWay:      item.ContactWay.String,
			Location:        item.Location.String,
			ViewCount:       int64(item.ViewCount),
			LikeCount:       int64(item.LikeCount),
			CommentCount:    int64(item.CommentCount),
			Status:          lostfound.LostFoundStatus(status),
			CreatedAt:       createTime,
			LostFoundTime:   lostFoundTime,
			PublisherName:   publisherName,
			PublisherAvatar: publisherAvatar,
		})
	}

	fmt.Printf("lostFoundList: %v", lostFoundList)
	return &lostfound.GetLostFoundListResponse{
		List:  lostFoundList,
		Total: int64(len(lostFoundList)),
		Page:  in.Page,
		Limit: in.Limit,
	}, nil
}
