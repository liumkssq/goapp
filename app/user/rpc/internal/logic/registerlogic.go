package logic

import (
	"context"
	"database/sql"

	"github.com/liumkssq/goapp/app/user/model"
	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"
	"github.com/liumkssq/goapp/pkg/tool"
	"github.com/liumkssq/goapp/pkg/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterError = xerr.NewMsg("user has been registered")
var ErrGenerateTokenError = xerr.NewMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewMsg("账号或密码不正确")

var ErrUserNotExistError = xerr.NewMsg("用户不存在")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 1. 查询用户是否已存在 (通过手机号)
	users, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	if users != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "该手机号已被注册 mobile:%s", in.Phone)
	}

	// 2. 验证短信验证码 (可选，从Redis获取验证码并校验)
	if in.VerificationCode != "" {
		// 实际项目中应该从Redis验证验证码
		// codeKey := fmt.Sprintf("sms:code:register:%s", in.Phone)
		// code, err := l.svcCtx.RedisClient.Get(codeKey)
		// if err != nil || code != in.VerificationCode {
		//    return nil, errors.Wrapf(xerr.NewErrMsg("验证码错误"), "验证码校验失败 phone:%s", in.Phone)
		// }
		// l.svcCtx.RedisClient.Del(codeKey) // 使用后删除验证码
	}

	// 3. 事务中创建用户
	var userId int64
	if err := l.svcCtx.UsersModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 3.1 创建用户记录
		u := &model.User{
			Username: in.Username,
			Phone:    in.Phone,
		}

		// 3.2 密码加密
		if len(in.Password) > 0 {
			genPasswd, err := tool.GenPasswordHash([]byte(in.Password))
			if err != nil {
				return errors.Wrapf(xerr.New(xerr.DB_ERROR, "ab"), "密码加密失败 err:%v", err)
			}
			u.PasswordHash = string(genPasswd)
		}

		// 3.3 设置默认值
		u.UserStatus = 1 // 正常状态

		// 3.4 添加新的学校相关信息
		if in.Campus != "" {
			u.Campus = sql.NullString{String: in.Campus, Valid: true}
		}
		if in.College != "" {
			u.College = sql.NullString{String: in.College, Valid: true}
		}
		if in.Major != "" {
			u.Major = sql.NullString{String: in.Major, Valid: true}
		}
		if in.EnrollmentYear > 0 {
			u.EnrollmentYear = sql.NullInt64{Int64: int64(in.EnrollmentYear), Valid: true}
		}
		if in.UserRole != "" {
			u.UserRole = sql.NullString{String: in.UserRole, Valid: true}
		}
		if in.StudentId != "" {
			u.StudentId = sql.NullString{String: in.StudentId, Valid: true}
		}

		// 3.5 插入数据库
		insertResult, err := l.svcCtx.UsersModel.Insert(ctx, session, u)
		if err != nil {
			return err
		}
		// 3.6 获取自增ID
		userId, err = insertResult.LastInsertId()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return &user.RegisterResponse{
		Username: in.Username,
		UserId:   userId,
	}, nil
}
