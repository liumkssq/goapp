package logic

import (
	"context"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/social/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/social/rpc/social"
	"github.com/liumkssq/goapp/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupListLogic) GroupList(in *social.GroupListReq) (*social.GroupListResp, error) {
	// 验证用户ID
	_, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewMsg("用户ID格式错误"), "parse uint error: %v", err)
	}

	// 获取用户所在的群组ID列表
	userIdStr := in.UserId // 使用原始字符串ID调用模型方法
	groupMembers, err := l.svcCtx.GroupMembersModel.ListByUserId(l.ctx, userIdStr)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "list group members error: %v, user_id: %s", err, userIdStr)
	}

	// 如果用户不在任何群组中，返回空结果
	if len(groupMembers) == 0 {
		return &social.GroupListResp{}, nil
	}

	// 提取群组ID并转换为字符串
	var groupIdsStr []string
	for _, member := range groupMembers {
		groupIdsStr = append(groupIdsStr, strconv.FormatUint(member.GroupId, 10))
	}

	// 获取群组详情
	groups, err := l.svcCtx.GroupsModel.ListByGroupIds(l.ctx, groupIdsStr)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "list groups error: %v, group_ids: %v", err, groupIdsStr)
	}

	// 构建返回数据
	var resp social.GroupListResp
	var groupList []*social.Groups
	for _, group := range groups {
		var item social.Groups
		err = copier.Copy(&item, group)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewMsg("复制结构体失败"), "copy error: %v", err)
		}
		item.Id = strconv.FormatUint(group.Id, 10)
		item.CreatorUid = strconv.FormatUint(group.CreatorId, 10)
		groupList = append(groupList, &item)
	}
	resp.List = groupList

	return &resp, nil
}
