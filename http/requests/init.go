package requests

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_t "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni *ut.UniversalTranslator
	validate *validator.Validate
	trans ut.Translator
)

func init() {
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	validate = binding.Validator.Engine().(*validator.Validate)
	zh_t.RegisterDefaultTranslations(validate, trans)
}

func Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}
