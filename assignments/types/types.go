package types

import (
	helpers "go-academy/helpers"
)

type User struct {
	firstName string
	surname   string
}

func (user *User) FirstName() string {
	return user.firstName
}

func (user *User) Surname() string {
	return user.surname
}

func (user *User) SetFirstName(str string) {
	user.firstName = helpers.RemoveNewlineFromStr(str)
}

func (user *User) SetSurname(str string) {
	user.surname = helpers.RemoveNewlineFromStr(str)
}
