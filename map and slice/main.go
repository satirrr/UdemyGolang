package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	//syntac to create a map
	//myMap := make(map[string]string)
	//or <dont use this>
	//var anotherMap map[string]string

	myMap := make(map[string]User)
	me := User{
		FirstName: "musa",
		LastName:  "mahatir",
	}

	myMap["me"] = me

	log.Println(myMap)
	log.Println(myMap["me"])
	log.Println(myMap["me"].FirstName)
	log.Println(myMap["me"].LastName)

	//TODO : explore this:
	//TODO : nextMap := make(map[string]interface{})

	slicePlayground()
	switchPlayground()
}
func slicePlayground() {
	sliceOfNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(sliceOfNumber)
	log.Println(sliceOfNumber[:5])
	log.Println(sliceOfNumber[5:])
}

func switchPlayground() {
	myVar := 6
	log.Println(myVar % 2)
	switch myVar / 2 {
	case 0:
		log.Println("genap")
	case 1:
		log.Println("ganjil")
	default:
		log.Println("what are you typing?")
	}
}
