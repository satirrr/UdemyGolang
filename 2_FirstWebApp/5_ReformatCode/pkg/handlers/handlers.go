package handlers

import (
	"github.com/satirrr/UdemyGolang/2_FirstWebApp/5_ReformatCode/pkg/renders"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	renders.NewRenderTemplate(writer, "home.page.tmpl")
}

func About(writer http.ResponseWriter, request *http.Request) {
	renders.NewRenderTemplate(writer, "about.page.tmpl")
}
