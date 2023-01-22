package main

import (
	"github.com/satirrr/UdemyGolang/2_FirstWebApp/5_ReformatCode/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("starting application on port : %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
