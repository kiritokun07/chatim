package httpresp

import (
	"fmt"
	"net/http"
	"strconv"

	"chatim/shared/errorx"

	"github.com/tealeg/xlsx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	codeServerError    = 1
	serviceUnavailable = "服务器竟然开小差，一会儿再试试吧"
)

type errResp struct {
	Code int    `json:"code"`
	Desc string `json:"desc,omitempty"`
}

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (e *errResp) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Desc
}

func Http(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	if err != nil {
		HttpErr(w, r, err)
	} else {
		HttpOkJson(w, r, data)
	}
}

func HttpErr(w http.ResponseWriter, r *http.Request, err error) {
	codeErr, ok := errorx.FromError(err)
	if ok {
		httpx.WriteJson(w, codeErr.Status(), errResp{
			Code: codeErr.Code(),
			Desc: codeErr.Error(),
		})
	} else {
		httpx.WriteJson(w, http.StatusInternalServerError, errResp{
			Code: codeServerError,
			Desc: serviceUnavailable,
		})
		logx.WithContext(r.Context()).Error(err)
	}
}

func HttpOkJson(w http.ResponseWriter, r *http.Request, data interface{}) {
	httpx.OkJson(w, &Body{
		Code: 0,
		Msg:  "ok",
		Data: data,
	})
}

func XlsxResponse(file *xlsx.File, name string, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "applicationnd.ms-excel")
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s.xlsx;filename*=utf-8''%s.xlsx", name, name))
	w.Header().Add("Cache-Control", "max-age=0")
	w.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	_ = file.Write(w)
}

func ImageResponse(png []byte, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(png)))
	_, _ = w.Write(png)
}
