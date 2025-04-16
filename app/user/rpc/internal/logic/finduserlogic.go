package logic

import (
	"context"
	"github.com/liumkssq/goapp/app/user/model"

	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找用户
func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	// todo: add your logic here and delete this line
	var (
		userEntities []*model.User
		err          error
	)
	if in.Phone != "" {
		userEntity, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.Phone)
		if err != nil {
			return nil, err
		}
		userEntities = append(userEntities, userEntity)
	} else if in.Name != "" {
		userEntity, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, in.Name)
		if err != nil {
			return nil, err
		}
		userEntities = append(userEntities, userEntity)
	} else if len(in.Ids) > 0 {
		userEntities, err = l.svcCtx.UsersModel.ListByIds(l.ctx, in.Ids)
		if err != nil {
			return nil, err
		}
	}
	var resp []*user.UserInfoResponse
	for _, userEntity := range userEntities {
		resp = append(resp, &user.UserInfoResponse{
			UserId:   int64(userEntity.Id),
			Username: userEntity.Username,
			Nickname: userEntity.Nickname.String,
			Phone:    userEntity.Phone,
		})
	}
	return &user.FindUserResp{
		User: resp,
	}, nil
}
