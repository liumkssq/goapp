package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserNotificationsModel = (*customUserNotificationsModel)(nil)

type (
	// UserNotificationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserNotificationsModel.
	UserNotificationsModel interface {
		userNotificationsModel
	}

	customUserNotificationsModel struct {
		*defaultUserNotificationsModel
	}
)

// NewUserNotificationsModel returns a model for the database table.
func NewUserNotificationsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserNotificationsModel {
	return &customUserNotificationsModel{
		defaultUserNotificationsModel: newUserNotificationsModel(conn, c, opts...),
	}
}
