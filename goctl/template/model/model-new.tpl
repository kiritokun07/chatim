
func new{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) *default{{.upperStartCamelObject}}Model {
	m := &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: sqlc.NewConn(conn, c){{else}}conn:conn{{end}},
		table:      {{.table}},
	}
	m.initBulkInserter(conn)
	return m
}
