package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LostFoundCommentModel = (*customLostFoundCommentModel)(nil)

type (
	// LostFoundCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLostFoundCommentModel.
	LostFoundCommentModel interface {
		lostFoundCommentModel
	}

	customLostFoundCommentModel struct {
		*defaultLostFoundCommentModel
	}
)

// NewLostFoundCommentModel returns a model for the database table.
func NewLostFoundCommentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LostFoundCommentModel {
	return &customLostFoundCommentModel{
		defaultLostFoundCommentModel: newLostFoundCommentModel(conn, c, opts...),
	}
}
