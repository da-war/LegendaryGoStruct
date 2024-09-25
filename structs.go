package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDay  string
	createdAt time.Time
}

func newUser(firstName, lastName, birthday string) (*user, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("First Name, LastName and birthDate is required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDay:  birthday,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := "Rana "
	lastName := "Dawar"
	birthDate := "2 December 1999"

	var appUser *user
	appUser, err := newUser(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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
