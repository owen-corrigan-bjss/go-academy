package main

import (
	"bufio"
	"fmt"
	types "go-academy/types"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	cmdLineReader := bufio.NewReader(os.Stdin)

	var user1 types.User
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
