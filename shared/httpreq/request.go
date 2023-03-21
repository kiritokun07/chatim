package httpreq

import (
	"errors"
	"net/http"
	"strings"

	"chatim/shared/validate"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var (
	xValidator = validate.NewValidator()
)

func Parse(r *http.Request, v interface{}) error {
	if err := httpx.Parse(r, v); err != nil {
		return err
	}
	if err := xValidator.Validator.Struct(v); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			transData := errs.Translate(xValidator.Trans)
			s := strings.Builder{}
			for _, v := range transData {
				s.WriteString(v)
				s.WriteString(" ")
			}
			return errors.New(s.String())
		}
		return err
	}
	return nil
}
