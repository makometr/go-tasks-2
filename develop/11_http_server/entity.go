package main

import (
	"context"
	"fmt"
	"sync"
)

// TODO rename file to module entity

type Event struct {
	ID     int
	UserID int
	Name   string
	Date   CalendarDay
}

type eventcontextKey int

const (
	name eventcontextKey = iota
	userID
	date
)

func newEvent(ID, userID int, name string, date CalendarDay) *Event {
	return &Event{ID: ID, UserID: userID, Name: name, Date: date}
}

type EventStorage struct {
	data   []Event
	lastID int
	rwm    sync.RWMutex
}

func newEventStorage(oldEvents []Event) *EventStorage {
	var maxID int
	for i := 0; i < len(oldEvents); i++ {
		if maxID < oldEvents[i].ID {
			maxID = oldEvents[i].ID
		}
	}
	return &EventStorage{data: oldEvents, lastID: maxID, rwm: sync.RWMutex{}}
}

func (es *EventStorage) getEvent(ID int) (Event, error) {
	es.rwm.RLock()
	defer es.rwm.RUnlock()

	index, err := es.findIndex(ID)
	if err != nil {
		return Event{}, nil
	}
	return es.data[index], nil
}

func (es *EventStorage) addEvent(e Event) error {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	if e.ID >= es.lastID {
		return fmt.Errorf("invalid ID")
	}
	es.data = append(es.data, e)
	es.lastID = e.ID

	return nil
}

func (es *EventStorage) updateEvent(ID int, ctx context.Context) error {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	index, err := es.findIndex(ID)
	if err != nil {
		return err
	}

	isChanged := false
	if newName, ok := ctx.Value(name).(string); ok {
		es.data[index].Name = newName
		isChanged = true
	}
	if newUserID, ok := ctx.Value(name).(int); ok {
		es.data[index].UserID = newUserID
		isChanged = true
	}
	if newDate, ok := ctx.Value(name).(CalendarDay); ok {
		es.data[index].Date = newDate
		isChanged = true
	}

	if isChanged == false {
		return fmt.Errorf("no changes was provided")
	}
	return nil
}

func (es *EventStorage) findIndex(ID int) (int, error) {
	for i := 0; i < len(es.data); i++ {
		if es.data[i].ID == ID {
			return i, nil
		}
	}
	return 0, fmt.Errorf("no event with sucn index")
}

// TODO check нужен ли мютекс

type CalendarDay struct {
	day  int
	week int
	year int
}
