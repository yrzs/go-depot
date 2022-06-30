package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

/**
声明 ValidError 相关的结构体和类型
标准库 errors 的 New 方法实现非常简单，errorString 是一个结构体，内含一个 s 字符串，也只有一个 Error 方法，就可以认定为 error 类型
在 Go 语言中，如果一个类型实现了某个 interface 中的所有方法，那么编译器就会认为该类型实现了此 interface，它们是”一样“的。
*/
type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(ctx *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := ctx.ShouldBind(v)
	if err != nil {
		v := ctx.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for k, v := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     k,
				Message: v,
			})
		}
		return false, errs
	}
	return true, nil
}
