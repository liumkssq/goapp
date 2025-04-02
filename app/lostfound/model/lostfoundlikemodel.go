package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LostFoundLikeModel = (*customLostFoundLikeModel)(nil)

type (
	// LostFoundLikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLostFoundLikeModel.
	LostFoundLikeModel interface {
		lostFoundLikeModel
	}

	customLostFoundLikeModel struct {
		*defaultLostFoundLikeModel
	}
)

// NewLostFoundLikeModel returns a model for the database table.
func NewLostFoundLikeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LostFoundLikeModel {
	return &customLostFoundLikeModel{
		defaultLostFoundLikeModel: newLostFoundLikeModel(conn, c, opts...),
	}
}
