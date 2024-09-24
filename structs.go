package main

import "fmt"

func main() {
	firstName := getUserData("Please Enter your first name:")
	lastName := getUserData("Please Enter your last name:")
	birthDate := getUserData("Please Enter your Date of birth")

	fmt.Print(firstName, lastName, birthDate)
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
