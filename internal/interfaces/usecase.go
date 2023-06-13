package interfaces

import "github.com/AndreySuhonosov/calendar/internal/domain"

type EventUseCase interface {
	AddEvent(event domain.Event) error
	DeleteEvent(eventId int) error
	ChangeEvent(NewEvent domain.Event) error
	EventList() ([]domain.Event, error)
}
