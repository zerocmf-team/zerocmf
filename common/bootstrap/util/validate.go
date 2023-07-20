package util

import "strings"

func IsStringEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
