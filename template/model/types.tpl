type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
		query     string
        queryArgs []interface{}
        limit     int
        offset    int
        orderBy   string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
