package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LostFoundItemModel = (*customLostFoundItemModel)(nil)

type (
	// LostFoundItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLostFoundItemModel.
	LostFoundItemModel interface {
		lostFoundItemModel
	}

	customLostFoundItemModel struct {
		*defaultLostFoundItemModel
	}
)

// NewLostFoundItemModel returns a model for the database table.
func NewLostFoundItemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LostFoundItemModel {
	return &customLostFoundItemModel{
		defaultLostFoundItemModel: newLostFoundItemModel(conn, c, opts...),
	}
}
