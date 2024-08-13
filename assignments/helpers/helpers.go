package helpers

import "strings"

func RemoveNewlineFromStr(str string) string {
	return strings.Replace(str, "\n", "", 1)
}
