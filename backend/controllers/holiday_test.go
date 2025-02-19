package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"holiday_calendar/config"
	"holiday_calendar/models"

	"github.com/gorilla/mux"
)

// Setup function to initialize the database connection
func setup() {
	config.ConnectDB()
}

// Test GetHolidays API
func TestGetHolidays(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/api/holidays", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListHolidays)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Optionally, check if the response contains valid JSON
	var holidays []models.Holiday
	err = json.Unmarshal(rr.Body.Bytes(), &holidays)
	if err != nil {
		t.Errorf("Response body is not valid JSON: %v", err)
	}
}

// Test AddHoliday API
func TestAddHoliday(t *testing.T) {
	setup()

	newHoliday := models.Holiday{
		Date: "2025-12-25",
		Name: "Test Holiday",
	}

	body, _ := json.Marshal(newHoliday)
	req, err := http.NewRequest("POST", "/api/holidays", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHoliday)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

// Test DeleteHoliday API
func TestDeleteHoliday(t *testing.T) {
	setup()

	// Add a holiday to delete
	newHoliday := models.Holiday{
		Date: "2025-12-31",
		Name: "Holiday to Delete",
	}
	body, _ := json.Marshal(newHoliday)
	addReq, _ := http.NewRequest("POST", "/api/holidays", bytes.NewBuffer(body))
	addReq.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	addRec := httptest.NewRecorder()

	// Call the AddHoliday handler
	AddHoliday(addRec, addReq)

	// Check if the holiday was added successfully
	if addRec.Code != http.StatusCreated {
		t.Fatalf("Failed to add holiday: got status %v", addRec.Code)
	}

	// Extract the added holiday's ID from the response
	var addedHoliday models.Holiday
	err := json.Unmarshal(addRec.Body.Bytes(), &addedHoliday)
	if err != nil {
		t.Fatalf("Failed to parse added holiday response: %v", err)
	}

	// Delete the holiday using the extracted ID
	deleteReq, _ := http.NewRequest("DELETE", "/api/holidays/"+addedHoliday.ID.Hex(), nil)
	deleteRec := httptest.NewRecorder()

	// Use mux to create a router and handle the request
	r := mux.NewRouter()
	r.HandleFunc("/api/holidays/{id}", DeleteHoliday)

	// Call the DeleteHoliday handler
	r.ServeHTTP(deleteRec, deleteReq)

	// Check if the deletion was successful
	if deleteRec.Code != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", deleteRec.Code, http.StatusNoContent)
	}
}
