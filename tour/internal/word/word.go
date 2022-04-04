package word

import (
	"strings"
	"unicode"
)

// ToUpper 字母全部转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 字母全部转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	//将所有的下划线替换成空格
	s = strings.Replace(s, "_", " ", -1)
	//将s中以空格分开的单词的首字母大写
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	// 只对首字母做小写
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
