package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDay  string
	createdAt time.Time
}

func newUser(firstName, lastName, birthday string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthDay:  birthday,
		createdAt: time.Now(),
	}
}

func main() {
	firstName := "Rana"
	lastName := "Dawar"
	birthDate := "2 December 1999"

	var appUser user
	appUser = newUser(firstName, lastName, birthDate)
	appUser.outPutUserDetails()
	appUser.updateUserName()
	appUser.outPutUserDetails()

	// var myUser user
	// myUser = user{
	// 	firstName: "daar",
	// 	lastName:  "ka",
	// 	birthDay:  "mahool",
	// }

	// myUser.outPutUserDetails()

}

func (myStruct user) outPutUserDetails() {

	fmt.Println("Touba", myStruct)
}

func (myStruct *user) updateUserName() {
	myStruct.firstName = "Dawarrr"
	myStruct.lastName = "Abdullahh"
}

func getUserData(str string) string {
	fmt.Println(str)
	var value string
	fmt.Scan(&value)
	return value
}
