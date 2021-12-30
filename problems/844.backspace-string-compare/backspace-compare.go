package backspace_compare

import (
	"strings"
)

func backspaceCompare(s string, t string) bool {
	if strings.Compare(cut(s), cut(t)) == 0 {
		return true
	}

	return false
}

func cut(str string) string {
	idx := strings.Index(str, "#")
	if idx != -1 {
		if idx == 0 {
			str = strings.Replace(str, "#", "", 1)
		} else {
			result := str[idx-1 : idx+1]
			str = strings.Replace(str, result, "", 1)
		}
		str = cut(str)
	}
	return str
}
