package string

import (
	"bytes"
	"strings"
	"unicode"
)

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

func toLowerSnakeCase(s string) string {
	var buf bytes.Buffer
	var lastWasUpper bool
	for i, r := range s {
		if unicode.IsUpper(r) && i != 0 && !lastWasUpper {
			buf.WriteRune('_')
		}
		lastWasUpper = unicode.IsUpper(r)
		buf.WriteRune(unicode.ToLower(r))
	}
	return buf.String()
}
