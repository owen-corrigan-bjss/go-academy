package helpers

import (
	"strings"
	"testing"
)

func TestRemovesSingleNewlineFromStr(test *testing.T) {
	initialString := "hello this is a string \n"
	formatedString := RemoveNewlineFromStr(initialString)

	if strings.Contains(formatedString, "\n") {
		test.Error("formated string still contains \n")
	}
}

func TestRemovesNewlineFromMiddleOfStr(test *testing.T) {
	initialString := "hello this \n is a string"
	formatedString := RemoveNewlineFromStr(initialString)

	if strings.Contains(formatedString, "\n") {
		test.Error("formated string still contains \n")
	}
}

func TestRemovesMultipleNewLinesFromStr(test *testing.T) {
	initialString := "hello this \n is a string\n"
	formatedString := RemoveNewlineFromStr(initialString)

	if strings.Contains(formatedString, "\n") {
		test.Error("formated string still contains \n")
	}

	initialString = "hello \n \n \nthis \n is a string\n"
	formatedString = RemoveNewlineFromStr(initialString)

	if strings.Contains(formatedString, "\n") {
		test.Error("formated string still contains \n")
	}
}
