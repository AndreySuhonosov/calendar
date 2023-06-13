package interfaces

import "github.com/AndreySuhonosov/calendar/internal/domain"

type EventRepository interface {
	AddEvent(event domain.Event) error
	GetEvent(eventId int) (domain.Event, error)
	DeleteEvent(eventId int) error
	ChangeEvent(NewEvent domain.Event) error
	EventList() ([]domain.Event, error)
}
