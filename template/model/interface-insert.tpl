Where(query string, args ...interface{}) *default{{.upperStartCamelObject}}Model
Limit(limit int) *default{{.upperStartCamelObject}}Model
Offset(offset int) *default{{.upperStartCamelObject}}Model
OrderBy(query string) *default{{.upperStartCamelObject}}Model
First(ctx context.Context) (*{{.upperStartCamelObject}}, error)
Find(ctx context.Context) ([]*{{.upperStartCamelObject}}, error)
Count(ctx context.Context) (int64, error)
Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)