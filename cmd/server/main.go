package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/db"
	"github.com/dan-nicholls/project-sourdough/internal/web"
)

func main() {
	fmt.Println("Starting Project Sourdough...")

	// Load Configurations
	port := 3000
	dbPath := "./orders.db"

	// Initialise DB
	sqlite := &db.Sqlite3{}
	if err := sqlite.Connect(dbPath); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer sqlite.Close()

	if err := sqlite.EnsureSchema(); err != nil {
		log.Fatalf("Failed to intialise schema: %v", err)
	}
	
	// Setup app
	appService := app.New(sqlite)

	server := &http.Server{
		Handler: web.NewRouter(appService),
		Addr: fmt.Sprintf(":%d", port),
	}

	fmt.Printf("🚀 Server running at port: %d\n", port)
	log.Fatal(server.ListenAndServe())
}
