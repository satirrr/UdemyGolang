package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "clark",
		"last_name": "abduloh",
		"has_dog": true
	},
	{
		"first_name": "alek",
		"last_name": "sukamok",
		"has_dog": false
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("error unmarshalling json", err)
	}

	log.Printf("unmarshalled json: %v", unmarshalled)

	//write a json from a slice

	var mySlice []Person

	var marshall1 Person
	marshall1.FirstName = "samsul"
	marshall1.LastName = "amdah"
	marshall1.HasDog = true

	marshall2 := Person{"asmod", "samson", false}

	mySlice = append(mySlice, marshall1, marshall2)

	newJson, err1 := json.MarshalIndent(mySlice, "", "   ")
	if err1 != nil {
		log.Println("error marshalling", err1)
	}

	fmt.Println(string(newJson))
}
