package main

import (
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(writer http.ResponseWriter, request *http.Request) {
	RenderTemplate(writer, "home.page.tmpl")
}

func About(writer http.ResponseWriter, request *http.Request) {
	RenderTemplate(writer, "about.page.tmpl")
}

//
//
func RenderTemplate(writer http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(writer, nil)
	if err != nil {
		log.Println("error parsing template :", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("starting application on port : %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
