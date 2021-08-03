package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// TODO rename file to module entity

type Event struct {
	ID     int
	UserID int
	Name   string
	Date   CalendarDay
}

func (e *Event) MarshalJSON() ([]byte, error) {
	buffer := struct {
		ID     int    `json:"id"`
		UserID int    `json:"user_id"`
		Name   string `json:"name"`
		Date   string `json:"date"`
	}{
		ID:     e.ID,
		UserID: e.UserID,
		Name:   e.Name,
		Date:   e.Date.String(),
	}
	return json.Marshal(&buffer)
}

func (e *Event) UnmarshalJSON(b []byte) error {
	var buffer struct {
		ID     int    `json:"id"`
		UserID int    `json:"user_id"`
		Name   string `json:"name"`
		Date   string `json:"date"`
	}
	err := json.Unmarshal(b, &buffer)
	if err != nil {
		return err
	}

	e.ID = buffer.ID
	e.UserID = buffer.UserID
	e.Name = buffer.Name
	e.Date = *NewCalendarDatFromString(buffer.Date)
	return nil
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

func (es *EventStorage) addEvent(data DataToAddNewEvent) error {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	es.data = append(es.data, Event{ID: es.getNewID(), UserID: data.UserID, Name: data.Name, Date: *NewCalendarDatFromString(data.Date)})

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

func (es *EventStorage) DeleteEvent(ID int) error {
	index, err := es.findIndex(ID)
	if err != nil {
		return err
	}

	es.data = append(es.data[:index], es.data[:index]...)
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

func (cd CalendarDay) String() string {
	return fmt.Sprintf("%d.%d.%d", cd.day, cd.week, cd.year)
}

func NewCalendarDatFromString(s string) *CalendarDay {
	numbers := strings.Split(s, ".")
	if len(numbers) != 3 {
		return &CalendarDay{}
	}

	day, err := strconv.Atoi(numbers[0])
	if err != nil {
		return &CalendarDay{}
	}
	week, err := strconv.Atoi(numbers[1])
	if err != nil {
		return &CalendarDay{}
	}
	year, err := strconv.Atoi(numbers[2])
	if err != nil {
		return &CalendarDay{}
	}

	var cd CalendarDay
	cd.day = day
	cd.week = week
	cd.year = year
	return &cd
}
