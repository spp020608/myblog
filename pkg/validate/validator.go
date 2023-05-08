package validate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// trans 翻译器定义
var trans ut.Translator

// InitTranslator 初始化翻译器
func InitTranslator(locale string) (err error) {

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil
	}
	// 修改 gin 框架中的 Validator 引擎属性，实现自定制

	// 相当于给我们的翻译器,注册了一个读取 json 标签的函数
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(zhT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	var found bool
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	trans, found = uni.GetTranslator(locale)
	if !found {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	if err != nil {
		return err
	}

	return
}

//
// TranslateError 翻译错误信息,
// 当错误类型为普通错误的时候, 返回值 isValidationErrors 为 false, commonError 为 err.Error(), validationErrorsMap 为 nil
//
// 当错误类型为验证库定义的错误类型 validator.ValidationErrors 的时候,
// 返回值 isValidationErrors 为 true, commonError 为 "", validationErrorsMap 为验证库返回的错误信息组成的 map
//
//	@parameter	err 通用错误信息
//
func TranslateError(err error) string {
	// 必须将原始错误信息转换成验证库自身的错误信息,否则无法翻译
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	m := removeTopStruct(errs.Translate(trans))

	var message string
	for key, value := range m {
		message += key + ":" + value + "; "
	}
	return message
}

//
// removeTopStruct 删除 Validator 库错误信息中的结构体名称, 返回给前端的结构体名称
//	@return	map[string]string 所有的错误信息
//
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}

	for field, err := range fields {
		key := field[strings.Index(field, ".")+1:]
		res[key] = err
	}
	return res

}
