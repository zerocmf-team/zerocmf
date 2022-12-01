package Required

import "strings"

func String(s string) (res bool) {
	if strings.TrimSpace(s) != "" {
		res = true
		return
	}
	res = false
	return
}
