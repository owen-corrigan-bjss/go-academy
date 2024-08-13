package helpers

import (
	"regexp"
)

func RemoveNewlineFromStr(str string) string {
	regex := regexp.MustCompile(`\n *`)
	return regex.ReplaceAllString(str, "")
}
