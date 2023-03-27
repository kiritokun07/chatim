
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{.keys}}
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return ret,err
}


func (m *default{{.upperStartCamelObject}}Model) BatchInsert(ctx context.Context,data *{{.upperStartCamelObject}}) error {
	return m.bulkInserter.Insert({{.expressionValues}})
}

func (m *default{{.upperStartCamelObject}}Model) Flush() {
	m.bulkInserter.Flush()
}

func (m *default{{.upperStartCamelObject}}Model) initBulkInserter(conn sqlx.SqlConn) {
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	inserter, err := sqlx.NewBulkInserter(conn, query)
	if err != nil {
		panic(err)
	}
	m.bulkInserter = inserter
}

