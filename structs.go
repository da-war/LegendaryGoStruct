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

func main() {
	firstName := "Rana"
	lastName := "Dawar"
	birthDate := "2 December 1999"

	var appUser user
	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDay:  birthDate,
		createdAt: time.Now(),
	}
	appUser.outPutUserDetails()

	var myUser user
	myUser = user{
		firstName: "daar",
		lastName:  "ka",
		birthDay:  "mahool",
	}

	myUser.outPutUserDetails()

}

func (myStruct user) outPutUserDetails() {

	fmt.Println("Touba", myStruct)
}

func getUserData(str string) string {
	fmt.Println(str)
	var value string
	fmt.Scan(&value)
	return value
}
