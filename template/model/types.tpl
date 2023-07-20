type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
		query     string
        queryArgs []interface{}
        limit     int32
        offset    int32
        orderBy   string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
