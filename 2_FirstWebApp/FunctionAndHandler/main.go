package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Home in order for func to handle a request, it has to have 2 param
func Home(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "HOMEPAGE")
	if err != nil {
		return
	}
}

func About(writer http.ResponseWriter, request *http.Request) {
	sum := AddValue(2, 2)
	_, err := fmt.Fprintf(writer, "ABOUT, calculate result : %d", sum)
	if err != nil {
		return
	}
}

func AddValue(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Sprintf("starting application on port : %s", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	log.Println(err)
}
