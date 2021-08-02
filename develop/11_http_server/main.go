package main

import (
	"config"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type application struct {
	events EventStorage
	db     DataManipulator
	server *http.Server
}

func (app *application) initDateStorage(cfg *config.Config) {
	app.db = newFileDB(cfg.DbFilename)
	evs, err := app.db.openConnection()
	if err != nil {
		log.Fatalf("File db error open %v", err)
	}
	app.events.initEventStorage(evs)
}

func (app *application) initHTTPServer(cfg *config.Config) {
	http.HandleFunc("/event", app.getEvent)
	http.HandleFunc("/create_event", app.AddEvent)
	http.HandleFunc("/update_event", app.UpdateEvent)
	http.HandleFunc("/delete_event", app.DeleteEvent)
	// http.HandleFunc("/events_for_day", app.Hello)
	// http.HandleFunc("/events_for_week", app.Hello)
	// http.HandleFunc("/events_for_month", app.Hello)
	app.server = &http.Server{Addr: cfg.Port, Handler: http.DefaultServeMux}

}

func (app *application) Shutdown() {
	err := app.db.closeConnction(app.events.data)
	if err != nil {
		log.Printf("Error while closing DB; %\nv", err)
	} else {
		log.Println("DB closed!")
	}

	err = app.server.Shutdown(context.Background())
	if err != nil {
		log.Printf("Error while closing server; %\nv", err)
	} else {
		log.Println("Server closed!")
	}
}

func main() {
	cfg := config.NewDefaultConfig()

	var app application
	app.initDateStorage(cfg)
	app.initHTTPServer(cfg)
	defer app.Shutdown()

	// fmt.Println(time.Time.ISOWeek(time.Now()))
	// fmt.Println(time.Time.YearDay(time.Now()))

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("server: started at address %v", cfg.Port)
		serverErrors <- http.ListenAndServe(":8080", nil)
	}()

	for {
		select {
		case err := <-serverErrors:
			fmt.Printf("server error %v\n", err)
		case <-osSignals:
			fmt.Println("service shutdown...")
			return
		}
	}
}
