package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/web/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(app *app.AppService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// API Routes
	r.With(TokenAuthMiddleware(app.TokenStore)).Post("/api/order", CreateOrderHandler(app))
	r.Post("/api/auth", AuthHandler(app.TokenStore, app.AccessCode))	

	// UI Routes
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/",  fs))

	r.Get("/", HomeHandler)

	r.Get("/order", func (w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.OrderView(app.FormOptions)).ServeHTTP(w, r)
	})
	
	r.Get("/success", SuccessHandler)
	r.Get("/error", ErrorHandler)

	return r
}
