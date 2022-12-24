package dao

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsAdminUserDao = (*customUmsAdminUserDao)(nil)

type (
	// UmsAdminUserDao is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminUserDao.
	UmsAdminUserDao interface {
		umsAdminUserDao
	}

	customUmsAdminUserDao struct {
		*defaultUmsAdminUserDao
	}
)

// NewUmsAdminUserDao returns a Dao for the database table.
func NewUmsAdminUserDao(conn sqlx.SqlConn, c cache.CacheConf) UmsAdminUserDao {
	return &customUmsAdminUserDao{
		defaultUmsAdminUserDao: newUmsAdminUserDao(conn, c),
	}
}
