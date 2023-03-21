package errorx

import (
	"errors"
	"fmt"

	"chatim/shared/errorx/grpcerrordetails"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type (
	CodeError interface {
		error
		Status() int
		Code() int
		Desc() string
	}

	codeError struct {
		status int
		code   int
		desc   string
	}
)

func (err *codeError) Desc() string {
	return err.desc
}

func (err *codeError) Error() string {
	return err.desc
}

func (err *codeError) Code() int {
	return err.code
}

func (err *codeError) Status() int {
	return err.status
}

func (err *codeError) String() string {
	return fmt.Sprintf("Status: %d, Code: %d, Desc: %s", err.status, err.code, err.desc)
}

func NewCodeError(code int, desc string) CodeError {
	return NewStatCodeError(400, code, desc)
}

func NewDefaultError(desc string) CodeError {
	return NewStatCodeError(400, 400000, desc)
}

func NewNotFoundError(desc string) CodeError {
	return NewCodeError(404, desc)
}

func NewStatCodeError(status, code int, desc string) CodeError {
	return &codeError{
		status: status,
		code:   code,
		desc:   desc,
	}
}

func FromError(err error) (codeErr CodeError, ok bool) {
	if err == nil {
		return nil, true
	}
	// CodeError
	var ce CodeError
	if ok := errors.As(err, &ce); ok {
		return ce, ok
	}
	// status.Status
	st, ok := status.FromError(err)
	if !ok {
		return nil, false
	}
	for _, d := range st.Details() {
		switch info := d.(type) {
		case *grpcerrordetails.GrpcErrorDetails:
			return NewCodeError(int(info.Code), info.Desc), true
		default:
			logx.Severef("Unexpected type: %+v", info)
		}
	}

	return nil, false
}
