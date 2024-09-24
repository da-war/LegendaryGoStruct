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
	// birthDate := getUserData("Please Enter your Date of birth")

	var appUser user
	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		// birthDay:birthDate,
		createdAt: time.Now(),
	}

	fmt.Println(appUser)
	if appUser.birthDay == "" {
		fmt.Println("Helllosh")
	}

}

func outPutUserDetails() {
	fmt.Println("")
}

func getUserData(str string) string {
	fmt.Println(str)
	var value string
	fmt.Scan(&value)
	return value
}
