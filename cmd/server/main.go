package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"

	"github.com/dan-nicholls/project-sourdough/internal/app"
	"github.com/dan-nicholls/project-sourdough/internal/db"
	"github.com/dan-nicholls/project-sourdough/internal/web"
)

func main() {
	fmt.Println("Starting Project Sourdough...")

	// Load Configurations
	var port = flag.Int("port", 3000, "Port to run the server on")
	var mock = flag.Bool("mock", false, "run with mock database")
	var code = flag.String("code", "Bread", "Code to allow auth access")

	flag.Parse()

	// Setup DB
	var database db.Database 
	if *mock {
		log.Println("Running in Mock Mode")
		database = &db.MockDB{}
		defer database.Close()
	} else {
		// Initialise DB
		dbPath := "./orders.db"

		database = &db.Sqlite3{}
		if err := database.Connect(dbPath); err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		defer database.Close()

		if err := db.InitialiseDB(database); err != nil {
			log.Fatalf("Failed to intialise schema: %v", err)
		}
	} 
	
	// Setup app
	appService := app.New(database, *code)

	server := &http.Server{
		Handler: web.NewRouter(appService),
		Addr: fmt.Sprintf(":%d", *port),
	}

	fmt.Printf("🚀 Server running at port: %d\n", *port)
	log.Fatal(server.ListenAndServe())
}
