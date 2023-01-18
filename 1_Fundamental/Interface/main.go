package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name      string
	Color     string
	Something int
}

func main() {
	dog := Dog{
		Name:  "herder",
		Breed: "strong",
	}
	PrintInfo(&dog)
}

func PrintInfo(a Animal) {
	fmt.Println("this animal says ", a.Says(), " and has ", a.NumberOfLegs(), " number of legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
