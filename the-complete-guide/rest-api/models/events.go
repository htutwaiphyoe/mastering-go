package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	DateTime    time.Time
	Location    string
	CreatedBy   int
}

func (e Event) Save() {
}

func Get() {}
