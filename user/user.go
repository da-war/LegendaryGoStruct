package user

import (
	"errors"
	"fmt"
	"os/user"
	"time"
)

type User struct {
	FirstName        string
	LastName         string
	BirthDateirthDay string
	CreatedAt        time.Time
}

func NewUser(firstName, lastName, birthday string) (*user.User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("First Name, LastName and birthDate is required")
	}
	return &user.User{
		firstName,
		astName,
		birthDate: birthday,
		CreatedAt: time.Now(),
	}, nil
}

func (myStruct user.User) OutPutUserDetails() {

	fmt.Println("Touba", myStruct)
}

func (myStruct *user.User) UpdateUserName() {
	myStruct.firstName = "Dawarrr"
	myStruct.lastName = "Abdullahh"
}

func GetUserData(str string) string {
	fmt.Println(str)
	var value string
	fmt.Scan(&value)
	return value
}
