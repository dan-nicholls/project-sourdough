package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/utils"
)

type OrderPayload struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Options struct {
		Shape string `json:"shape"`
		Flour string `json:"flour"`
		Topping string `json:"topping"`
		Scoring string `json:"scoring"`
	}
	Code string `json:"code"`
}

func (p OrderPayload) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Email == "" || !utils.IsValidEmail(p.Email) {
		return errors.New("valid email is required")
	}

	if p.Options.Shape == "" {
		return errors.New("Shape option is required")
	}

	if p.Options.Flour == "" {
		return errors.New("Flour option is required")
	}

	if p.Options.Topping == "" {
		return errors.New("Topping option is required")
	}

	if p.Options.Scoring == "" {
		return errors.New("Scoring option is required")
	}

	return nil
}

// POST - /order
func CreateOrderHandler(a *app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p OrderPayload
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			log.Printf("Error parsing form: %v", err)
			w.Header().Set("HX-Redirect", "/error")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return

		}

		// Basic Code Auth
		if p.Code != a.AccessCode {
			log.Printf("Unauthorized: code mismatch")
			w.Header().Set("HX-Redirect", "/error")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate payload
		if err := p.Validate(); err != nil {
			log.Printf("Error validating order: %v", err)
			w.Header().Set("HX-Redirect", "/error")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		order := app.Order{
			Name: p.Name,
			Email: p.Email, 
			Options: app.BreadOptions{
				Shape:  p.Options.Shape, 
				Flour:   p.Options.Flour,
				Topping: p.Options.Topping,
				Scoring: p.Options.Scoring,
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
			w.Header().Set("HX-Redirect", "/error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if t, ok := GetTokenFromContext(r.Context()); ok {
			a.TokenStore.MarkUsed(t.ID)
		}

		w.Header().Set("HX-Redirect", "/success")
		w.WriteHeader(http.StatusOK)
	}
}

// POST - /auth
type TokenResponse struct {
	Token string `json:"token"`
}

type AuthRequest struct {
	Code string `json:"code"`
}

func AuthHandler(store app.TokenStore, keyword string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ar AuthRequest
		err := json.NewDecoder(r.Body).Decode(&ar)
		if err != nil {
			log.Printf("Unable to parse AuthRequest", err)	
			w.Header().Set("HX-Redirect", "/error")
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		if ar.Code != keyword {
			log.Printf("Invalid keyword used: %v", ar.Code)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		t := app.Token{
			ID: uuid.New().String(),
			IssuedAt: time.Now(),
			ExpiresAt: time.Now().Add(time.Minute*5),
			Used: false,
		}

		store.SaveToken(t)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TokenResponse{Token: t.ID})
	}
}
