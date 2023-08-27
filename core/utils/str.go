package utils

import "strings"

func TrimPath(path string) string {
	return strings.Trim(path, "/")
}

func EqualsPath(a, b string) bool {
	return strings.ToUpper(TrimPath(a)) == strings.ToUpper(TrimPath(b))
}
