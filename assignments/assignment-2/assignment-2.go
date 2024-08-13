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

func removeNewLine(str string) string {
	return strings.Replace(str, "\n", "", -1)
}

func (user *User) FirstName() string {
	return user.firstName
}

func (user *User) Surname() string {
	return user.surname
}

func (user *User) SetFirstName(str string) {
	user.firstName = removeNewLine(str)
}

func (user *User) SetSurname(str string) {
	user.surname = removeNewLine(str)
}

func inputFormatter(str string) string {
	return strings.Replace(str, "\n", "", 1)
}

func main() {
	cmdLineReader := bufio.NewReader(os.Stdin)

	var user1 User
	var name string
	var err error

	pl("Please input your first name")
	name, err = cmdLineReader.ReadString(('\n'))

	if err != nil {
		log.Fatal("something went wrong")
	}

	user1.SetFirstName(name)

	pl("please input your surname")

	name, err = cmdLineReader.ReadString(('\n'))

	if err != nil {
		log.Fatal("something went wrong")
	}

	user1.SetSurname(name)

	fmt.Printf("Hello %s %s nice to meet you\n", user1.FirstName(), user1.Surname())
}
