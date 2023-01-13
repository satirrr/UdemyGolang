package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	exampleOfSlice := []string{createSlice(), "king of spade"}
	fmt.Println(exampleOfSlice)

	exampleOfSlice = append(exampleOfSlice, "ace of spade")
	fmt.Println("this is the new exampleOfSlice ", exampleOfSlice)
	//append is creating a new array, not really add a new data

	for i, aSlice := range exampleOfSlice {
		fmt.Println(i, aSlice)
	}
}

func createSlice() string {
	return "five of diamond"
}
