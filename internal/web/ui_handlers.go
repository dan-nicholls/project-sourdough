package web

import (
	"net/http"

	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/web/templates"
	"github.com/a-h/templ"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(templates.Home("Project Sourdough")).ServeHTTP(w, r)
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(templates.OrderSuccess()).ServeHTTP(w, r)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(templates.OrderError()).ServeHTTP(w, r)
}

func OrderViewHandler(app *app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderView(app.FormOptions)).ServeHTTP(w, r)
	}
}
