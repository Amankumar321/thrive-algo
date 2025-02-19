package routes

import (
	"holiday_calendar/controllers"

	"github.com/gorilla/mux"
)

func RegisterHolidayRoutes(r *mux.Router) {
	r.HandleFunc("/api/holidays", controllers.ListHolidays).Methods("GET")
	r.HandleFunc("/api/holidays", controllers.AddHoliday).Methods("POST")
	r.HandleFunc("/api/holidays/{id}", controllers.DeleteHoliday).Methods("DELETE")
}
