package {{.PkgName}}

import (
	"net/http"

	{{if .HasRequest}}"chatim/shared/errorx"{{end}}
	{{if .HasRequest}}"chatim/shared/httpreq"{{end}}
	"chatim/shared/httpresp"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpreq.Parse(r, &req); err != nil {
			httpresp.HttpErr(w, r, errorx.NewStatCodeError(http.StatusNotAcceptable, 2, err.Error()))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}httpresp.Http(w, r, resp, err){{else}}httpresp.Http(w, r, nil, err){{end}}
	}
}
