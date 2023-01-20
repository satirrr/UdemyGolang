package renders

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(writer http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(writer, nil)
	if err != nil {
		log.Println("error parsing template :", err)
	}
}
