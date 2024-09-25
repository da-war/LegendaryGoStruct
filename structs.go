package main

import (
	"fmt"

	"dawar.com/user"
)

func main() {
	firstName := "dawar"
	lastName := "Dawar"
	birthDate := "2 December 1999"
	var appUser *user.User
	appUser, err := user.New(firstName, lastName, birthDate)
	appAdmin, err := user.NewAdmin("dawar@gmail.com", "Dawar@123", firstName, lastName, birthDate)

	newAdmin, err := user.MakeUserAdmin("emailcom", "dawar123", appUser)

	fmt.Println("BirthDay:", appUser)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	appUser.OutPutUserDetails()
	appAdmin.OutPutAdmin()
	newAdmin.OutPutAdmin()

	// appUser.UpdateUserName()
	// appUser.OutPutUserDetails()

}
