func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}

// 根据条件进行查询一条数据
func (m *default{{.upperStartCamelObject}}Model) First(ctx context.Context) (*{{.upperStartCamelObject}}, error) {
	query := m.query

	queryArgs := m.queryArgs
	orderBy := m.orderBy
	var resp {{.upperStartCamelObject}}
	sql := fmt.Sprintf("select %s from %s", {{.lowerStartCamelObject}}Rows, m.table)

	if query != "" {
		sql += " where " + query
	}

	// 排序
    if orderBy != "" {
        sql += fmt.Sprintf(" ORDER BY %s", orderBy)
    }

	sql += " AND deleted_at = 0 limit 1"

	err := m.QueryRowNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 根据条件进行列表查询
func (m *default{{.upperStartCamelObject}}Model) Find(ctx context.Context) ([]*{{.upperStartCamelObject}}, error) {

	query := m.query
	queryArgs := m.queryArgs
	orderBy := m.orderBy

	var resp []*{{.upperStartCamelObject}}
	sql := fmt.Sprintf("select %s from %s", {{.lowerStartCamelObject}}Rows, m.table)

	if query != "" {
    	sql += " where " + query + " AND deleted_at = 0"
    }else {
        sql += " where deleted_at = 0"
    }

	// 排序
	if orderBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", orderBy)
	}

	limit := m.limit
	offset := m.offset

	// 查询条件
	if limit > 0 {
		sql += fmt.Sprintf(" LIMIT %d", limit)
	}

	if offset > 0 {
		sql += fmt.Sprintf(" OFFSET %d", offset)
	}

	err := m.QueryRowsNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 统计字段
func (m *default{{.upperStartCamelObject}}Model) Count(ctx context.Context) (int64, error) {
	query := m.query
	queryArgs := m.queryArgs
	sql := fmt.Sprintf("select count({{.originalPrimaryKey}}) from %s", m.table)
	if query != "" {
		sql += " where " + query
	}
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}