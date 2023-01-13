package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

//create a function for user struct
func (user *User) printFirstName() string {
	return user.FirstName
}

func main() {
	user := User{
		FirstName: "alek",
		LastName:  "subarjo",
	}

	var saipul User
	saipul.FirstName = "samsat"

	log.Println(saipul.FirstName)

	log.Println(user, "birthDate : ", user.BirthDate)

	log.Println("print using user func ", saipul.printFirstName())
	
}
