package main

import (
	"fmt"

	"dawar.com/user"
)

func main() {
	firstName := "Rana "
	lastName := "Dawar"
	birthDate := "2 December 1999"

	var appUser *user.User
	appUser = &user.User{
		FirstName:        firstName,
		LastName:         lastName,
		BirthDateirthDay: birthDate,
	}

	fmt.Println("BirthDay:", appUser)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// appUser.outPutUserDetails()
	// appUser.updateUserName()
	// appUser.outPutUserDetails()

}
