package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) getEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get event\n")
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
	fmt.Println("Data after addEvent:", app.events.data)

	fmt.Println(t)

}

func (app *application) UpdateEvent(w http.ResponseWriter, r *http.Request) {

}

func (app *application) DeleteEvent(w http.ResponseWriter, r *http.Request) {

}
