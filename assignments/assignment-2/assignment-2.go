package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var pl = fmt.Println

type User struct {
	firstName string
	surname   string
}

func (user *User) FormatName() {
	removeNewLine := func(str string) string {
		return strings.Replace(str, "\n", "", -1)
	}

	user.firstName = removeNewLine(user.firstName)
	user.surname = removeNewLine(user.surname)
}

func inputFormatter(str string) string {
	return strings.Replace(str, "\n", "", -1)
}

func main() {
	cmdLineReader := bufio.NewReader(os.Stdin)

	var user1 User
	var err error

	pl("Please input your first name")
	user1.firstName, err = cmdLineReader.ReadString(('\n'))

	if err != nil {
		log.Fatal("something went wrong")
	}

	pl("please input your surname")

	user1.surname, err = cmdLineReader.ReadString(('\n'))

	if err != nil {
		log.Fatal("something went wrong")
	}

	user1.FormatName()

	fmt.Printf("Hello %s %s nice to meet you\n", user1.firstName, user1.surname)
}
