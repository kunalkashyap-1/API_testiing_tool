package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kunalkashyap-1/API_testiing_tool/internal/routes"
	"github.com/kunalkashyap-1/API_testiing_tool/pkg/database"
	"github.com/kunalkashyap-1/API_testiing_tool/pkg/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatalf("Databse URL environment variable is required.")
	}
	db, err := database.Initialize(connStr)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close()

	//Initialise the router
	r := mux.NewRouter()
	//Apply logging middleware to the router
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)

	api := r.PathPrefix("/api").Subrouter()
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"Up and running, chief"}`))
	}).Methods("GET")

	// Register routes from different modules
	routes.RegisterAuthRoutes(api, db)    // Authentication routes
	routes.RegisterSenarioRoutes(api, db) // Scenario management routes
	routes.RegisterTestRoutes(api, db)    // Test execution and management routes
	routes.RegisterMetricRoutes(api, db)  // Metrics gathering and display routes
	routes.RegisterReportRoutes(api, db)  // Report generation routes
	routes.RegisterUserRoutes(api, db)    // User management routes

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %v...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
