package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"
const x1 = 100.0
const y1 = 0.0

// Home in order for func to handle a request, it has to have 2 param
func Home(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "HOMEPAGE")
	if err != nil {
		return
	}
}

func About(writer http.ResponseWriter, request *http.Request) {
	sum := addValue(2, 2)
	_, err := fmt.Fprintf(writer, "ABOUT, calculate result : %d", sum)
	if err != nil {
		return
	}
}

func addValue(x, y int) int {
	return x + y
}

func Divide(writer http.ResponseWriter, request *http.Request) {
	result, err := divideValue(x1, y1)
	if err != nil {
		fmt.Fprintf(writer, "cannot divide by zero")
		return
	}
	fmt.Fprintf(writer, fmt.Sprintf("%f divided by %f is equal to %f", x1, y1, result))
}

func divideValue(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divided by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Sprintf("starting application on port : %s", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	log.Println(err)
}
