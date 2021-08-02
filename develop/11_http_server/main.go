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

type app struct {
	db dataEventManipulator
}

func initServer() *app {
	// return &app{db: newFileDB("sas")}
	return nil
}

func main() {
	config.SayHello()
	cfg := config.NewDefaultConfig()
	db := newFileDB(cfg.DbFilename)
	if err := db.openConnection(); err != nil {
		log.Fatalf("File db errror open %v", err)
	}
	defer db.closeConnction()

	// app := initServer()
	// http.HandleFunc("/create_event", api.AddEvent)
	// http.HandleFunc("/events_for_day", api.Hello)
	// http.HandleFunc("/events_for_week", api.Hello)
	// http.HandleFunc("/events_for_month", api.Hello)
	server := &http.Server{Addr: cfg.Port, Handler: nil}

	// fmt.Println(time.Time.ISOWeek(time.Now()))
	// fmt.Println(time.Time.YearDay(time.Now()))

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	serverErrors := make(chan error, 1)
	go func() {
		// log.Infof("server: started at address %v", config)
		serverErrors <- http.ListenAndServe(":8080", nil)
	}()

	for {
		select {
		case err := <-serverErrors:
			fmt.Printf("server error %v\n", err)
		case <-osSignals:
			fmt.Println("service shutdown...")
			server.Shutdown(context.Background())
			return
		}
	}
}
