package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/a-h/templ"

	"github.com/dan-nicholls/project-sourdough/internal/web/templates"
)

type App struct {
	server http.Server
	port int
}

func (a *App) Start() error {
	http.Handle("/", templ.Handler(templates.Home("This is an intro text")))

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static", fs)

	fmt.Printf("Serving on port %d\n", a.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), nil)
	return err
}

func main() {
	fmt.Println("Hello World!")

	app := App{
		port: 3000,
	}

	log.Fatal(app.Start())
}
