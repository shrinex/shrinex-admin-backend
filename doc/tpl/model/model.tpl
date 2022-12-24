package {{.pkg}}
{{if .withCache}}
import (
	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{else}}

import (
	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{end}}
var _ {{.upperStartCamelObject}}Dao = (*custom{{.upperStartCamelObject}}Dao)(nil)

type (
	// {{.upperStartCamelObject}}Dao is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Dao.
	{{.upperStartCamelObject}}Dao interface {
		{{.lowerStartCamelObject}}Dao
	}

	custom{{.upperStartCamelObject}}Dao struct {
		*default{{.upperStartCamelObject}}Dao
	}
)

// New{{.upperStartCamelObject}}Dao returns a Dao for the database table.
func New{{.upperStartCamelObject}}Dao(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Dao {
	return &custom{{.upperStartCamelObject}}Dao{
		default{{.upperStartCamelObject}}Dao: new{{.upperStartCamelObject}}Dao(conn{{if .withCache}}, c{{end}}),
	}
}
