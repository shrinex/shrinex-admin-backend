
func new{{.upperStartCamelObject}}Dao(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) *default{{.upperStartCamelObject}}Dao {
	return &default{{.upperStartCamelObject}}Dao{
		{{if .withCache}}CachedConn: sqlc.NewConn(conn, c){{else}}conn:conn{{end}},
		table:      {{.table}},
	}
}
