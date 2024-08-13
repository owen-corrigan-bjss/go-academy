package main

import (
	"bufio"
	"fmt"
	helpers "go-academy/helpers"
	"log"
	"os"
)

var pl = fmt.Println

type User struct {
	firstName string
	surname   string
}

func (user *User) SetFirstName(str string) {
	user.firstName = helpers.RemoveNewlineFromStr(str)
}

func (user *User) SetSurname(str string) {
	user.surname = helpers.RemoveNewlineFromStr(str)
}

func main() {
	cmdLineReader := bufio.NewReader(os.Stdin)

	var user1 User

	for len(user1.firstName) == 0 {
		pl("Please input your first name")
		name, err := cmdLineReader.ReadString(('\n'))
		if err != nil {
			log.Fatal(err)
		}
		name = helpers.RemoveNewlineFromStr(name)
		user1.SetFirstName(name)
	}

	for len(user1.surname) == 0 {
		pl("Please input your surname")
		name, err := cmdLineReader.ReadString(('\n'))
		if err != nil {
			log.Fatal(err)
		}
		name = helpers.RemoveNewlineFromStr(name)
		user1.SetSurname(name)
	}

	fmt.Printf("Hello %s %s nice to meet you\n", user1.firstName, user1.surname)
}
