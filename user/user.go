package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First Name, LastName and birthDate is required")
	}
	hello := time.Now()
	return &User{
		firstName,
		lastName,
		birthDate,
		hello,
	}, nil
}

func (myStruct User) OutPutUserDetails() {

	fmt.Println("Touba", myStruct)
}

func (myStruct *User) UpdateUserName() {
	myStruct.firstName = "Dawarrr"
	myStruct.lastName = "Abdullahh"
}

func GetUserData(str string) string {
	fmt.Println(str)
	var value string
	fmt.Scan(&value)
	return value
}
