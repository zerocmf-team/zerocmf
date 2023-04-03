package test

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"testing"
)

type User struct {
	FirstName string     `json:"first_name" validate:"required" label:"姓"`
	LastName  string     `json:"last_name" validate:"required" label:"名"`
	Age       uint8      `validate:"gte=0,lte=130"`
	Email     string     `validate:"required,email"`
	Addresses []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func translateOverride(trans ut.Translator) {

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}不能为空!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	type User struct {
		Username string `label:"用户名" validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}

func TestValidate(t *testing.T) {

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName: "",
		LastName:  "",
		Age:       135,
		Email:     "Badger.Smith@gmail.com",
		Addresses: []*Address{address},
	}

	zh := zh.New()
	uni = ut.New(zh, zh)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("zh")

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

	translateOverride(trans)

	err := validate.Struct(user)

	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		errs.Translate(trans)
	}

	return

	if err != nil {
		fmt.Println("=== error msg ====")
		fmt.Println(err)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		fmt.Println("\r\n=========== error field info ====================")
		//for _, err := range err.(validator.ValidationErrors) {
		//	// 列出效验出错字段的信息
		//	fmt.Println("Namespace: ", err.Namespace())
		//	fmt.Println("Fild: ", err.Field())
		//	fmt.Println("StructNamespace: ", err.StructNamespace())
		//	fmt.Println("StructField: ", err.StructField())
		//	fmt.Println("Tag: ", err.Tag())
		//	fmt.Println("ActualTag: ", err.ActualTag())
		//	fmt.Println("Kind: ", err.Kind())
		//	fmt.Println("Type: ", err.Type())
		//	fmt.Println("Value: ", err.Value())
		//	fmt.Println("Param: ", err.Param())
		//	fmt.Println()
		//}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
