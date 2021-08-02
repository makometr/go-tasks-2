package main

type dataEventManipulator interface {
	openConnection() error
	closeConnction() error
	readEvents() ([]Event, error)
	writeEvents([]Event) error

	addEvent(eventQueryOptions) error
	updateEvent(eventQueryOptions) (bool, error)
	deleteEvent(ID int) (bool, error)
	findEvents(eventQueryOptions) (Event, error)
}

type eventQueryOptions struct {
	ID     *int
	userID *int
	date   *CalendarDay
	name   *string
}

func findEventByID(ID int) eventQueryOptions {
	return eventQueryOptions{ID: &ID, userID: nil, date: nil}
}

func findByUserID(userID int) eventQueryOptions {
	return eventQueryOptions{ID: nil, userID: &userID, date: nil}
}

func findByDate(date CalendarDay) eventQueryOptions {
	return eventQueryOptions{ID: nil, userID: nil, date: &date}
}

func findByUserIDandDate(userID int, date CalendarDay) eventQueryOptions {
	return eventQueryOptions{ID: nil, userID: &userID, date: &date}
}
