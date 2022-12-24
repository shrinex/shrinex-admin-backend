
type (
	{{.lowerStartCamelObject}}Dao interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Dao struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
