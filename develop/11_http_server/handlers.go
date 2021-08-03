package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) getEvent(w http.ResponseWriter, r *http.Request) {
	eventID, ok := r.URL.Query()["id"]
	if !ok {
		encoder := json.NewEncoder(w)
		encoder.Encode(app.events.data)
		return
	}
	fmt.Fprintf(w, "get event %v\n", eventID)
}

type DataToAddNewEvent struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
}

func (app *application) AddEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t DataToAddNewEvent
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	app.events.addEvent(t)
}

type DateToUpdateEvent struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
}

func (app *application) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t DateToUpdateEvent
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	index, err := app.events.findIndex(t.ID)
	if err != nil {
		panic(err)
	}
	app.events.data[index] = Event{ID: t.ID, Name: t.Name, UserID: t.UserID, Date: CalendarDay{}}
}

type DataToDeleteEvent struct {
	ID int `json:"id"`
}

func (app *application) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t DataToDeleteEvent
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	err = app.events.DeleteEvent(t.ID)
	if err != nil {
		panic(err)
	}
}
