package user

import (
	"global-greet-client/apiconsumer"
	"global-greet-client/simpleio"
)

type User struct {
	Name   string
	Native string
}

func MakeUser() User {
	name := simpleio.UserInput("Enter your name: ")
	lang := simpleio.UserInput("What is your language? ")
	newUser := User{Name: name, Native: lang}
	return newUser
}

func (u *User) Greet(apiResponse []byte) string {
	greeting := apiconsumer.ParseToJSON(apiResponse)
	return greeting.Hello + ", " + u.Name + "! " + greeting.Welcome + "."
}
