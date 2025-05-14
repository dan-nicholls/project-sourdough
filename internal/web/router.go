package web

import (
	"net/http"
	"time"
	"encoding/json"

	"github.com/a-h/templ"
	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/web/templates"
)

func CreateOrderHandler(a *app.AppService) http.HandlerFunc {
	return func (w http.ResponseWriter, res *http.Request) {
		if res.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		err := res.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		options := app.BreadOptions{
			Shape: res.FormValue("shape"),
			Flour: res.FormValue("flour"),
			Topping: res.FormValue("topping"),
			Scoring: res.FormValue("scoring"),
		}

		order := app.Order{
			Name: res.FormValue("name"),
			Email: res.FormValue("email"),
			Options: options,
			Status: "pending",
			OrderDate: time.Now(),
		}
		
		err = order.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Submit order here
		result, err := a.InsertOrder(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}



func NewRouter(a *app.AppService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(templates.Home("This is an intro text")))
	mux.Handle("/order", templ.Handler(templates.Order(a.Config)))
	mux.HandleFunc("/api/v1/order", CreateOrderHandler(a))

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static", fs)

	return mux
}
