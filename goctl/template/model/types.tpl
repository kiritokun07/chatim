
type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
		bulkInserter *sqlx.BulkInserter
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
