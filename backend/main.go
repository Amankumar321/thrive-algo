package main

import (
	"log"
	"net/http"

	"holiday_calendar/config"
	"holiday_calendar/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RegisterHolidayRoutes(r)

	// Add CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", corsHandler.Handler(r))
}
