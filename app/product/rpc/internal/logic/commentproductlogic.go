package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	model "github.com/liumkssq/goapp/app/product/model"
	"github.com/liumkssq/goapp/app/product/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/product/rpc/product"
	userModel "github.com/liumkssq/goapp/app/user/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CommentProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentProductLogic {
	return &CommentProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论商品
func (l *CommentProductLogic) CommentProduct(in *product.CommentProductRequest) (*product.CommentProductResponse, error) {
	if in.GetUserId() == 0 {
		return nil, errors.New("无效的用户ID")
	}
	if in.GetId() == 0 {
		return nil, errors.New("无效的商品ID")
	}
	if in.GetContent() == "" {
		return nil, errors.New("评论内容不能为空")
	}

	var newCommentId int64

	// 使用事务确保原子性
	// IMPORTANT: Assumes CommentModel has a Trans method implemented in commentmodel.go
	err := l.svcCtx.CommentModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 1. 获取用户信息
		userInfo, err := l.svcCtx.UserModel.FindOne(ctx, uint64(in.GetUserId()))
		if err != nil {
			if err == userModel.ErrNotFound { // Check against user model's ErrNotFound
				l.Errorf("用户 %d 不存在", in.GetUserId())
				return errors.New("用户不存在")
			}
			l.Errorf("获取用户信息失败: %v", err)
			return errors.New("获取用户信息失败")
		}

		// 2. 创建评论数据
		commentData := &model.Comment{
			ProductId: uint64(in.GetId()),
			UserId:    uint64(in.GetUserId()),
			UserName:  userInfo.Nickname.String, // Use .String for sql.NullString
			UserAvatar: sql.NullString{
				String: userInfo.AvatarUrl.String, // Assuming field is AvatarUrl (sql.NullString)
				Valid:  userInfo.AvatarUrl.Valid,  // Assuming field is AvatarUrl (sql.NullString)
			},
			Content:   in.GetContent(),
			CreatedAt: time.Now(),
		}

		// 3. 插入评论 (使用事务 session)
		// Use raw field names from model
		commentFieldNamesExpectAutoSet := "`product_id`, `user_id`, `user_name`, `user_avatar`, `content`, `created_at`" // Ensure field names match commentmodel_gen.go
		// TEMPORARY WORKAROUND: Accessing underlying table name directly. FIX model interface!
		commentTableName := "`comment`" // Directly use table name, replace with l.svcCtx.CommentModel.TableName() after fixing model
		queryInsert := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", commentTableName, commentFieldNamesExpectAutoSet)
		result, err := session.ExecCtx(ctx, queryInsert,
			commentData.ProductId,
			commentData.UserId,
			commentData.UserName,
			commentData.UserAvatar,
			commentData.Content,
			commentData.CreatedAt, // Pass CreatedAt explicitly
		)
		if err != nil {
			l.Errorf("插入评论失败 (sql: %s): %v", queryInsert, err)
			return errors.New("评论失败")
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			l.Errorf("获取评论插入ID失败: %v", err)
			return errors.New("评论记录ID获取失败")
		}
		newCommentId = lastId

		// 4. 更新商品评论数 (使用事务 session)
		// TEMPORARY WORKAROUND: Accessing underlying table name directly. FIX model interface!
		productTableName := "`product`" // Directly use table name, replace with l.svcCtx.ProductModel.TableName() after fixing model
		queryUpdate := fmt.Sprintf("update %s set `comment_count` = `comment_count` + 1 where `id` = ?", productTableName)
		_, err = session.ExecCtx(ctx, queryUpdate, commentData.ProductId)
		if err != nil {
			l.Errorf("更新商品评论数失败 (sql: %s): %v", queryUpdate, err)
			return errors.New("评论计数更新失败")
		}

		return nil // 事务成功
	})

	if err != nil {
		// 事务执行出错
		return nil, err
	}

	// 事务成功
	return &product.CommentProductResponse{
		CommentId: newCommentId,
		Message:   "评论成功",
	}, nil
}
