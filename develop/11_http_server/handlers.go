package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// func getEvent(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "get event\n")
// }

// func (app *app) AddEvent(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)
// 	var t struct {
// 		UserID int `json:"user_id"`
// 		Name string `json:"name"`
// 		date string `json:"date"`
// 	}
// 	err := decoder.Decode(&t)
// 	if err != nil {
// 		panic(err)
// 	}

// }
