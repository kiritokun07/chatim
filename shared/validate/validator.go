package validate

import (
	"log"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translational "github.com/go-playground/validator/v10/translations/zh"
)

type Validator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

func NewValidator() *Validator {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	err := translational.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatal(err)
	}
	return &Validator{
		Validator: validate,
		Trans:     trans,
	}
}
