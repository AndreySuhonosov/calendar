package domain

import "time"

type Event struct {
	Id          int
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}
