package splitstring

import "strings"

// 切割字符串
// example:
// abc, b=>[a c]
func Split(str string, sep string) []string {
	var ret []string
	index := strings.Index(str, sep)
	for index > 0 {
		ret = append(ret, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	return ret
}
