func new{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf, opts ...cache.Option{{end}}) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: sqlc.NewConn(conn, c, opts...){{else}}conn:conn{{end}},
		table:      {{.table}},
	}
}

func (m *default{{.upperStartCamelObject}}Model) withSession(session sqlx.Session) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn:m.CachedConn.WithSession(session){{else}}conn:sqlx.NewSqlConnFromSession(session){{end}},
		table:      {{.table}},
	}
}

func (m *default{{.upperStartCamelObject}}Model) Where(query string, args ...interface{}) *default{{.upperStartCamelObject}}Model {
	m.query = query
	m.queryArgs = args
	return m
}

func (m *default{{.upperStartCamelObject}}Model) Limit(limit int) *default{{.upperStartCamelObject}}Model {
	m.limit = limit
	return m
}

func (m *default{{.upperStartCamelObject}}Model) Offset(offset int) *default{{.upperStartCamelObject}}Model {
	m.offset = offset
	return m
}

func (m *default{{.upperStartCamelObject}}Model) OrderBy(orderBy string) *default{{.upperStartCamelObject}}Model {
	m.orderBy = orderBy
	return m
}