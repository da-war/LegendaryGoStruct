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

type Admin struct {
	email     string
	password  string
	user      User
	createdAt time.Time
}

// func NewAdmin(email, password string) (Admin, error) {
// 	if email == "" || password == "" {
// 		return nil, errors.New("Email and Password are required")
// 	}

// 	return Admin{
// 		email,
// 		password,
// 		user: User{
// 			firstName: "Dawar",
// 			lastName:  "Filana",
// 			birthDate: "Dimkana",
// 			createdAt: time.Now(),
// 		},
// 	}, nil

// }
func NewAdmin(email, password, firstName, lastName, birthDate string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and Password are required")
	}
	user, err := New(firstName, lastName, birthDate)
	if err != nil {
		return nil, err
	}
	return &Admin{
		email:     email,
		password:  password,
		user:      *user,
		createdAt: time.Now(),
	}, nil
}

func MakeUserAdmin(email string, password string, user *User) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and Password are required")
	}
	hello := time.Now()
	return &Admin{
			email,
			password,
			*user,
			hello,
		},
		nil
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
func (myStruct Admin) OutPutAdmin() {
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
