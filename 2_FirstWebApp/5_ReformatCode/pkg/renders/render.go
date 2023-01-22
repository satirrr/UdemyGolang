package renders

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(writer http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(writer, nil)
	if err != nil {
		log.Println("error parsing template :", err)
	}
}

var templateCache = make(map[string]*template.Template)

func NewRenderTemplate(writter http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template i our cache
	_, inMap := templateCache[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and add it to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template
		log.Println("using cached template")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(writter, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[t] = tmpl

	return nil
}
