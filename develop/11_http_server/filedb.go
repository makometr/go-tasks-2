package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Concrete file DB realiztion

type fileDB struct {
	filename string
	filePtr  *os.File

	data   []Event
	lastID int
}

func newFileDB(filename string) *fileDB {
	return &fileDB{filename: filename, filePtr: nil, data: []Event{}}
}

func (db *fileDB) genNextID() int {
	defer func() {
		db.lastID++
	}()
	return db.lastID
}

func (db *fileDB) openConnection() error {
	ptr, err := os.OpenFile(db.filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		var event Event
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			db.data = []Event{}
			return err
		}
		db.data = append(db.data, event)
		if event.ID > db.lastID {
			db.lastID = event.ID
		}
	}

	db.filePtr = ptr
	return nil
}

func (db *fileDB) closeConnction() error {
	defer db.filePtr.Close()
	if err := db.filePtr.Truncate(0); err != nil {
		return err
	}
	db.filePtr.Seek(0, 0)

	for i := 0; i < len(db.data); i++ {
		jsonData, err := json.Marshal(db.data[i])
		if err != nil {
			return err
		}
		db.filePtr.Write(jsonData)
	}
	return nil
}

func (db *fileDB) addEvent(e eventQueryOptions) error {
	if e.date == nil || e.name == nil || e.userID == nil {
		return fmt.Errorf("no event data provided")
	}
	if e.ID == nil {
		newID := db.genNextID()
		e.ID = &newID
	}
	db.data = append(db.data, Event{ID: *e.ID, userID: *e.userID, name: *e.name, date: *e.date})
	return nil
}

func (db *fileDB) updateEvent(e eventQueryOptions) (bool, error) {
	if e.ID == nil {
		return false, fmt.Errorf("id should be provided")
	}
	if e.name == nil && e.userID == nil && e.date == nil {
		return false, fmt.Errorf("at least one filed should be provided")
	}

	for i := 0; i < len(db.data); i++ {
		if db.data[i].ID == *e.ID {
			if e.date != nil {
				db.data[i].date = *e.date
			}
			if e.userID != nil {
				db.data[i].userID = *e.userID
			}
			if e.name != nil {
				db.data[i].userID = *e.userID
			}
			return true, nil
		}
	}
	return false, nil
}

func (db *fileDB) deleteEvent(ID int) (bool, error) {
	var i int
	for i = 0; i < len(db.data); i++ {
		if db.data[i].ID == ID {
			break
		}
	}
	if i == len(db.data) {
		return false, nil
	}

	db.data = append(db.data[:i], db.data[i+1:]...)
	return true, nil
}

func (db *fileDB) findEvents(e eventQueryOptions) ([]Event, error) {
	if e.ID != nil {
		for i := 0; i < len(db.data); i++ {
			if db.data[i].ID == *e.ID {
				return []Event{db.data[i]}, nil
			}
		}
	}
	return nil, fmt.Errorf("error: todo")
}
