package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

// NewError 根据code和msg创建一个新的Error对象
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

// Error 返回错误对应的string格式信息
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d,错误信息：%s", e.code, e.msg)
}

// Code 返回错误的错误码
func (e *Error) Code() int {
	return e.code
}

// Msg 返回错误的简要信息
func (e *Error) Msg() string {
	return e.msg
}

// Msgf 格式化方式返回错误的简要信息
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

// Details 返回错误的details
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 为Error对象增加details信息
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

// StatusCode 根据我们定义的错误码返回其对应的http状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
