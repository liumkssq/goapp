// 社交服务群入群请求处理逻辑

package logic

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/liumkssq/goapp/app/social/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/social/rpc/social"
	"github.com/liumkssq/goapp/app/social/socialmodels"
	"github.com/liumkssq/goapp/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 错误定义
var (
	ErrNoGroupRequest  = xerr.NewMsg("no group request")
	ErrAlreadyHandled  = xerr.NewMsg("already handled")
	ErrInvalidOperator = xerr.NewMsg("invalid operator")
)

type GroupPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInHandleLogic {
	return &GroupPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutInHandleLogic) GroupPutInHandle(in *social.GroupPutInHandleReq) (*social.GroupPutInHandleResp, error) {
	// 1. 查询请求记录
	req, err := l.svcCtx.GroupRequestsModel.FindOne(l.ctx, int64(in.GroupReqId))
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, ErrNoGroupRequest
	}

	// 2. 查询操作者权限
	gm, err := l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx,
		in.HandleUid, uint64ToString(req.GroupId))
	if err != nil {
		return nil, err
	}
	if gm.RoleLevel < 1 {
		return nil, ErrInvalidOperator
	}

	// 3. 判断是否已经处理过
	if req.HandleResult.Valid {
		return nil, ErrAlreadyHandled
	}

	// 4. 开启事务，处理请求
	err = l.svcCtx.GroupRequestsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 更新处理状态
		req.HandleUserId = sql.NullInt64{Valid: true, Int64: stringToInt64(in.HandleUid)}
		req.HandleTime = sql.NullTime{Valid: true, Time: time.Now()}
		req.HandleResult = sql.NullInt64{Valid: true, Int64: int64(in.HandleResult)}

		// 更新请求记录
		if err := l.svcCtx.GroupRequestsModel.Update(ctx, session, req); err != nil {
			return err
		}

		// 如果同意，添加成员
		if in.HandleResult == 1 {
			// 创建群组成员记录
			member := &socialmodels.GroupMembers{
				GroupId:    req.GroupId,
				UserId:     req.ReqUserId,
				RoleLevel:  0, // 普通成员
				JoinTime:   time.Now(),
				JoinSource: sql.NullInt64{Valid: req.JoinSource.Valid, Int64: req.JoinSource.Int64},
				InviterId:  sql.NullInt64{Valid: req.InviterId.Valid, Int64: req.InviterId.Int64},
				OperatorId: sql.NullInt64{Valid: true, Int64: stringToInt64(in.HandleUid)},
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
				DelState:   0,
			}

			_, err := l.svcCtx.GroupMembersModel.Insert(ctx, session, member)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &social.GroupPutInHandleResp{}, nil
}

func stringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// 辅助函数：将uint64转换为字符串
