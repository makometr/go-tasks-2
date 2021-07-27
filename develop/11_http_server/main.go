package main

import (
	"fmt"
	"net/http"
)

type app struct {
	db dataEventManipulator
}

func initServer() *app {
	// return &app{db: newFileDB("sas")}
	return nil
}

func main() {
	// app := initServer()
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
	// fmt.Println(time.Time.ISOWeek(time.Now()))
	// fmt.Println(time.Time.YearDay(time.Now()))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
