package logic

import (
	"context"
	"database/sql"

	"github.com/liumkssq/goapp/app/social/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/social/rpc/social"
	"github.com/liumkssq/goapp/app/social/socialmodels"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/liumkssq/goapp/pkg/wuid"
	"github.com/liumkssq/goapp/pkg/xerr"

	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群要求
func (l *GroupCreateLogic) GroupCreate(in *social.GroupCreateReq) (*social.GroupCreateResp, error) {
	// todo: add your logic here and delete this line

	groupId, err := strconv.ParseUint(wuid.GenUid(l.svcCtx.Config.Mysql.DataSource), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "parse groupId err %v", err)
	}

	creatorId, err := strconv.ParseUint(in.CreatorUid, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "parse creatorId err %v", err)
	}

	groups := &socialmodels.Groups{
		Id:        groupId,
		Name:      in.Name,
		Icon:      sql.NullString{String: in.Icon, Valid: true},
		CreatorId: creatorId,
		//IsVerify:   true,
		IsVerify: 0, // 0表示不需要验证
	}

	err = l.svcCtx.GroupsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		_, err := l.svcCtx.GroupsModel.Insert(l.ctx, session, groups)

		if err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "insert group err %v req %v", err, in)
		}

		_, err = l.svcCtx.GroupMembersModel.Insert(l.ctx, session, &socialmodels.GroupMembers{
			GroupId:   groups.Id,
			UserId:    creatorId,
			RoleLevel: int64(constants.CreatorGroupRoleLevel),
		})
		if err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "insert group member err %v req %v", err, in)
		}
		return nil
	})

	time.Sleep(2 * time.Second)

	return &social.GroupCreateResp{
		Id: strconv.FormatUint(groups.Id, 10),
	}, err
}
