package split_string

import "strings"

// Split ... 切割字符串

// Split ...
func Split(str string, sep string) []string {
	// str: "abcdbf"       sep: "b"
	var ret []string
	index := strings.Index(str, sep)
	for index > 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
