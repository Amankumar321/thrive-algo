package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"holiday_calendar/config"
	"holiday_calendar/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListHolidays(w http.ResponseWriter, r *http.Request) {
	var holidayCollection = config.DB.Collection("holidays")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var holidays []models.Holiday
	cursor, err := holidayCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var holiday models.Holiday
		cursor.Decode(&holiday)
		holidays = append(holidays, holiday)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(holidays)
}

func AddHoliday(w http.ResponseWriter, r *http.Request) {
	var holidayCollection = config.DB.Collection("holidays")
	var holiday models.Holiday
	_ = json.NewDecoder(r.Body).Decode(&holiday)

	holiday.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := holidayCollection.InsertOne(ctx, holiday)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(holiday)
}

func DeleteHoliday(w http.ResponseWriter, r *http.Request) {
	var holidayCollection = config.DB.Collection("holidays")
	params := mux.Vars(r)
	id := params["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	res, err := holidayCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		http.Error(w, "Holiday not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
