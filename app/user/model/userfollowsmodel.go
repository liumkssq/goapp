package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserFollowsModel = (*customUserFollowsModel)(nil)

type (
	// UserFollowsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFollowsModel.
	UserFollowsModel interface {
		userFollowsModel
	}

	customUserFollowsModel struct {
		*defaultUserFollowsModel
	}
)

// NewUserFollowsModel returns a model for the database table.
func NewUserFollowsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserFollowsModel {
	return &customUserFollowsModel{
		defaultUserFollowsModel: newUserFollowsModel(conn, c, opts...),
	}
}
