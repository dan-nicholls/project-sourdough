package web

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/web/templates"
	"github.com/go-chi/chi/v5"
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
			w.Header().Set("HX-Redirect", "/order/failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := a.InsertOrder(order); err != nil {
			log.Printf("Error inserting order: %v", err)	
			w.Header().Set("HX-Redirect", "/order/failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Redirect", "/order/complete")
		w.WriteHeader(http.StatusOK)
	}
}

func CreateStepHandler(app *app.AppService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		stepIndexStr := chi.URLParam(r, "stepIndex")

		// Decode previous state from JSON
		var state map[string]string = make(map[string]string)
		
		err := r.ParseForm(); if err != nil {
			http.Error(w, "Error parsing form: "+err.Error(), http.StatusInternalServerError)
		}

		if len(r.Form) != 0 {
			// Get new form vals
			for i, k := range r.Form {
				if len(k) != 0 {
					state[i] = k[0]
				}
			}
		}

		if stepIndexStr == "start" {
			templates.OrderStart().Render(r.Context(), w)
			return
		}

		if stepIndexStr == "details" {
			templates.OrderDetails(state).Render(r.Context(), w)
			return
		}

		if stepIndexStr == "review" {
			templates.OrderReview(state).Render(r.Context(), w)
			return
		}

		stepIndex, err := strconv.Atoi(stepIndexStr)
		if err != nil {
			templates.OrderFailed().Render(r.Context(), w)
			return
		}

		if stepIndex >= len(app.Steps) {
			templates.OrderDetails(state).Render(r.Context(), w)
			return
		}

		// Save new answer
		step := app.Steps[stepIndex]
		err = templates.StepForm(step, stepIndex, len(app.Steps), state).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "render error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}


func NewRouter(app *app.AppService) http.Handler {
	r := chi.NewRouter()

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/",  fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.Home("This is an intro test")).ServeHTTP(w, r)
	})
	r.Post("/api/order", CreateOrderHandler(app))

	r.Get("/order", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderFlow()).ServeHTTP(w, r)
	})

	r.Get("/order/start", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderStart()).ServeHTTP(w, r)
	})

	r.Get("/order/complete", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderComplete()).ServeHTTP(w, r)
	})

	r.Get("/order/failed", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderFailed()).ServeHTTP(w, r)
	})

	r.Get("/order/step/{stepIndex}", CreateStepHandler(app))
	r.Post("/order/step/{stepIndex}", CreateStepHandler(app))
	
	return r
}
