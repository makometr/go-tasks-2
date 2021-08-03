package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// TODO rename file to module entity

type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
}

type eventcontextKey int

const (
	name eventcontextKey = iota
	userID
	date
)

func newEvent(ID, userID int, name string, date string) *Event {
	return &Event{ID: ID, UserID: userID, Name: name, Date: date}
}

func (e Event) toJSON() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) fromJSON(b []byte) error {
	return json.Unmarshal(b, e)
}

type EventStorage struct {
	data   []Event
	lastID int
	rwm    sync.RWMutex
}

func (es *EventStorage) getNewID() int {
	defer func() {
		es.lastID++
	}()
	return es.lastID
}

func (es *EventStorage) initEventStorage(oldEvents []Event) {
	for i := 0; i < len(oldEvents); i++ {
		if es.lastID < oldEvents[i].ID {
			es.lastID = oldEvents[i].ID
		}
	}
	es.data = oldEvents
	es.lastID = 1
	es.rwm = sync.RWMutex{}
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

func (es *EventStorage) addEvent(data DataToAddNewEvent) (int, error) {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	newID := es.getNewID()
	es.data = append(es.data, Event{ID: newID, UserID: data.UserID, Name: data.Name, Date: data.Date})

	return newID, nil
}

func (es *EventStorage) updateEvent(ID int, ctx context.Context) (Event, error) {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	index, err := es.findIndex(ID)
	if err != nil {
		return Event{}, err
	}

	if newName, ok := ctx.Value(name).(string); ok {
		es.data[index].Name = newName
	}
	if newUserID, ok := ctx.Value(userID).(int); ok {
		es.data[index].UserID = newUserID
	}
	if newDate, ok := ctx.Value(date).(string); ok {
		es.data[index].Date = newDate
	}

	return es.data[index], nil
}

func (es *EventStorage) DeleteEvent(ID int) (Event, error) {
	index, err := es.findIndex(ID)
	if err != nil {
		return Event{}, err
	}

	deletedEvent := es.data[index]
	if index != len(es.data)-1 {
		es.data[index] = es.data[len(es.data)-1]
	}
	es.data = es.data[:len(es.data)-1]

	return deletedEvent, nil
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

// func NewCalendarDatFromString(s string) *CalendarDay {
// 	numbers := strings.Split(s, ".")
// 	if len(numbers) != 3 {
// 		return &CalendarDay{}
// 	}

// 	day, err := strconv.Atoi(numbers[0])
// 	if err != nil {
// 		return &CalendarDay{}
// 	}
// 	week, err := strconv.Atoi(numbers[1])
// 	if err != nil {
// 		return &CalendarDay{}
// 	}
// 	year, err := strconv.Atoi(numbers[2])
// 	if err != nil {
// 		return &CalendarDay{}
// 	}

// 	var cd CalendarDay
// 	cd.day = day
// 	cd.week = week
// 	cd.year = year
// 	return &cd
// }
