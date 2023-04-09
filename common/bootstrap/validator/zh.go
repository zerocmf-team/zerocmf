package validator

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type Zh struct{}

func (rest *Zh) Validator() (trans ut.Translator, validate *validator.Validate) {
	zh := zh.New()
	uni := ut.New(zh, zh)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	//通过label标签返回自定义错误内容
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	zhTranslations.RegisterDefaultTranslations(validate, trans)
	translateOverride(trans, validate)
	return
}

func translateOverride(trans ut.Translator, validate *validator.Validate) {
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}不能为空!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}
