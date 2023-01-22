package handlers

import (
	"github.com/satirrr/UdemyGolang/2_FirstWebApp/5_ReformatCode/pkg/renders"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	renders.RenderTemplate(writer, "home.page.tmpl")
}

func About(writer http.ResponseWriter, request *http.Request) {
	renders.RenderTemplate(writer, "about.page.tmpl")
}
