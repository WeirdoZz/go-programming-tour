package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}
type ValidErrors []*ValidError

// Error 将错误打出
func (v *ValidError) Error() string {
	return v.Message
}

// Error 将所有错误连接，以逗号分割
func (v *ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

// Errors 将每个错误单独打出
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// BindAndValid 对入参校验进行二次封装，如果校验错误，调用Translator对其错误信息进行翻译
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	// 尝试进行参数绑定和入参校验
	err := c.ShouldBind(v)
	// 如果失败
	if err != nil {
		// 从ctx中获取translator
		v := c.Value("trans")
		// 将translator转成其本身的格式
		trans, _ := v.(ut.Translator)
		// 尝试将错误转成ValidationErrors
		verrs, ok := err.(validator.ValidationErrors)
		// 如果无法转的话说明没有validerror
		if !ok {
			return false, errs
		}

		// 对validerrors进行翻译然后返回
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
