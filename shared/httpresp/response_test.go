package httpresp

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"chatim/shared/errorx"

	"github.com/stretchr/testify/assert"
)

type message struct {
	Name string `json:"name"`
}

type tracedResponseWriter struct {
	headers     map[string][]string
	builder     strings.Builder
	hasBody     bool
	code        int
	lessWritten bool
	timeout     bool
}

func (w *tracedResponseWriter) Header() http.Header {
	return w.headers
}

func (w *tracedResponseWriter) Write(bytes []byte) (n int, err error) {
	if w.timeout {
		return 0, http.ErrHandlerTimeout
	}

	n, err = w.builder.Write(bytes)
	if w.lessWritten {
		n -= 1
	}
	w.hasBody = true

	return
}

func (w *tracedResponseWriter) WriteHeader(code int) {
	w.code = code
}

func TestHttpOkJson(t *testing.T) {
	w := tracedResponseWriter{
		headers: make(map[string][]string),
	}
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	HttpOkJson(&w, r, message{Name: "hello"})
	assert.Equal(t, http.StatusOK, w.code)
	assert.Equal(t, `{"name":"hello"}`, w.builder.String())
}

func TestHttpErr(t *testing.T) {
	w := tracedResponseWriter{
		headers: make(map[string][]string),
	}
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	HttpErr(&w, r, errors.New("system error"))
	assert.Equal(t, http.StatusInternalServerError, w.code)
	assert.Equal(t, `{"code":1,"desc":"服务器竟然开小差，一会儿再试试吧"}`, w.builder.String())

	w = tracedResponseWriter{
		headers: make(map[string][]string),
	}
	err := errorx.NewStatCodeError(406, 101, "customer error")
	HttpErr(&w, r, err)
	assert.Equal(t, http.StatusNotAcceptable, w.code)
	assert.Equal(t, `{"code":101,"desc":"customer error"}`, w.builder.String())

}
