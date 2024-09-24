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

	outPutUserDetails(&appUser)

}

func outPutUserDetails(myStruct *user) {
	userO := *myStruct
	fmt.Println("heheh", userO)
}

func getUserData(str string) string {
	fmt.Println(str)
	var value string
	fmt.Scan(&value)
	return value
}
