package web

import (
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/web/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateOrderHandler(a *app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %v", err)
			w.Header().Set("HX-Redirect", "/order/failed")
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		order := app.Order{
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
			Options: app.BreadOptions{
				Shape:   r.FormValue("shape"),
				Flour:   r.FormValue("flour"),
				Topping: r.FormValue("topping"),
				Scoring: r.FormValue("scoring"),
			},
			Status:    "pending",
			OrderDate: time.Now(),
		}

		if err := order.Validate(); err != nil {
			log.Printf("Error validating order: %v", err)	
			w.Header().Set("HX-Redirect", "/error")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := a.InsertOrder(order); err != nil {
			log.Printf("Error inserting order: %v", err)	
			w.Header().Set("HX-Redirect", "/order/failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Redirect", "/success")
		w.WriteHeader(http.StatusOK)
	}
}

func NewRouter(app *app.AppService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// API Routes
	r.Post("/api/order", CreateOrderHandler(app))

	// UI Routes
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/",  fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.Home("This is an intro test")).ServeHTTP(w, r)
	})

	r.Get("/order", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderView(app.FormOptions)).ServeHTTP(w, r)
	})
	
	r.Get("/success", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderComplete()).ServeHTTP(w,r)
	})

	r.Get("/error", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderError()).ServeHTTP(w, r)
	})

	return r
}
