package convert

import "strconv"

type StrTo string

// String 返回自身的string形式
func (s StrTo) String() string {
	return string(s)
}

// Int 返回自身的Int形式和一个错误
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt 强制转换成Int形式
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// UInt32 返回自身的Uint32类型
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
