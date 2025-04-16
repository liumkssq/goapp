package logic

import (
	"context"
	"database/sql"

	"github.com/liumkssq/goapp/app/social/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/social/rpc/social"
	"github.com/liumkssq/goapp/app/social/socialmodels"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/liumkssq/goapp/pkg/xerr"

	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinLogic {
	return &GroupPutinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutinLogic) GroupPutin(in *social.GroupPutinReq) (*social.GroupPutinResp, error) {
	// todo: add your logic here and delete this line

	//  1. 普通用户申请 ： 如果群无验证直接进入
	//  2. 群成员邀请： 如果群无验证直接进入
	//  3. 群管理员/群创建者邀请：直接进入群

	var (
		inviteGroupMember *socialmodels.GroupMembers
		userGroupMember   *socialmodels.GroupMembers
		groupInfo         *socialmodels.Groups

		err error
	)

	reqUserId, err := strconv.ParseUint(in.ReqId, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "parse reqUserId err %v", err)
	}

	groupId, err := strconv.ParseUint(in.GroupId, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "parse groupId err %v", err)
	}

	// 将uint64转换为string以适配模型方法
	reqUserIdStr := in.ReqId
	groupIdStr := in.GroupId

	userGroupMember, err = l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx, reqUserIdStr, groupIdStr)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group member by groud id and  req id err %v, req %v, %v", err,
			in.GroupId, in.ReqId)
	}
	if userGroupMember != nil {
		return &social.GroupPutinResp{}, nil
	}

	groupReq, err := l.svcCtx.GroupRequestsModel.FindByGroupIdAndReqId(l.ctx, groupIdStr, reqUserIdStr)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group req by groud id and user id err %v, req %v, %v", err,
			in.GroupId, in.ReqId)
	}
	if groupReq != nil {
		return &social.GroupPutinResp{}, nil
	}

	inviterUserId := uint64(0)
	if in.InviterUid != "" {
		inviterUserId, err = strconv.ParseUint(in.InviterUid, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "parse inviterUserId err %v", err)
		}
	}

	groupReq = &socialmodels.GroupRequests{
		ReqUserId: reqUserId,
		GroupId:   groupId,
		ReqMsg: sql.NullString{
			String: in.ReqMsg,
			Valid:  true,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		JoinSource: sql.NullInt64{
			Int64: int64(in.JoinSource),
			Valid: true,
		},
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	}

	if inviterUserId > 0 {
		groupReq.HandleUserId = sql.NullInt64{
			Int64: int64(inviterUserId),
			Valid: true,
		}
	}

	createGroupMember := func() {
		if err != nil {
			return
		}
		err = l.createGroupMember(in, reqUserId, groupId, inviterUserId)
	}

	// 将 groupId 转换为 int64 适配 FindOne 方法
	groupInfo, err = l.svcCtx.GroupsModel.FindOne(l.ctx, strconv.FormatUint(groupId, 10))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group by groud id err %v, req %v", err, in.GroupId)
	}

	// 验证是否要验证
	if groupInfo.IsVerify == 0 {
		// 不需要
		defer createGroupMember()

		groupReq.HandleResult = sql.NullInt64{
			Int64: int64(constants.PassHandlerResult),
			Valid: true,
		}

		return l.createGroupReq(groupReq, true)
	}

	// 验证进群方式
	if constants.GroupJoinSource(in.JoinSource) == constants.PutInGroupJoinSource {
		// 申请
		return l.createGroupReq(groupReq, false)
	}

	inviterUserId, err = strconv.ParseUint(in.InviterUid, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "parse inviterUserId err %v", err)
	}

	inviterUserIdStr := in.InviterUid
	inviteGroupMember, err = l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx, inviterUserIdStr, groupIdStr)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group member by groud id and user id err %v, req %v",
			in.InviterUid, in.GroupId)
	}

	if constants.GroupRoleLevel(inviteGroupMember.RoleLevel) == constants.CreatorGroupRoleLevel || constants.
		GroupRoleLevel(inviteGroupMember.RoleLevel) == constants.ManagerGroupRoleLevel {
		// 是管理者或创建者邀请
		defer createGroupMember()

		groupReq.HandleResult = sql.NullInt64{
			Int64: int64(constants.PassHandlerResult),
			Valid: true,
		}
		return l.createGroupReq(groupReq, true)
	}
	return l.createGroupReq(groupReq, false)
}

func (l *GroupPutinLogic) createGroupReq(groupReq *socialmodels.GroupRequests, isPass bool) (*social.GroupPutinResp, error) {
	_, err := l.svcCtx.GroupRequestsModel.Insert(l.ctx, groupReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert group req err %v req %v", err, groupReq)
	}

	if isPass {
		return &social.GroupPutinResp{GroupId: strconv.FormatUint(groupReq.GroupId, 10)}, nil
	}

	return &social.GroupPutinResp{}, nil
}

func (l *GroupPutinLogic) createGroupMember(in *social.GroupPutinReq, userId, groupId, operatorId uint64) error {
	joinTime := time.Now()
	var joinSource sql.NullInt64
	var inviterId sql.NullInt64
	var operatorIdDb sql.NullInt64

	if in.JoinSource > 0 {
		joinSource = sql.NullInt64{Int64: int64(in.JoinSource), Valid: true}
	}

	if operatorId > 0 {
		operatorIdDb = sql.NullInt64{Int64: int64(operatorId), Valid: true}
		inviterId = sql.NullInt64{Int64: int64(operatorId), Valid: true}
	}

	groupMember := &socialmodels.GroupMembers{
		GroupId:    groupId,
		UserId:     userId,
		RoleLevel:  int64(constants.AtLargeGroupRoleLevel),
		JoinTime:   joinTime,
		JoinSource: joinSource,
		InviterId:  inviterId,
		OperatorId: operatorIdDb,
	}
	_, err := l.svcCtx.GroupMembersModel.Insert(l.ctx, nil, groupMember)
	if err != nil {
		return errors.Wrapf(xerr.NewDBErr(), "insert friend err %v req %v", err, groupMember)
	}

	return nil
}
