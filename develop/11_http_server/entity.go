package main

type Event struct {
	ID     int
	userID int
	name   string
	date   CalendarDay
}

func newEvent(ID, userID int, name string, date CalendarDay) *Event {
	return &Event{ID: ID, userID: userID, name: name, date: date}
}

type CalendarDay struct {
	day  int
	week int
	year int
}
