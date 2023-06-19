package local

import (
	"github.com/AndreySuhonosov/calendar/internal/domain"
	"github.com/AndreySuhonosov/calendar/internal/errors"
	"sort"
	"time"
)

type Event struct {
	Id          int
	Name        string
	Description string
	StartTime   int64
	EndTime     int64
}

type events []Event

func (e events) Len() int {
	return len(e)
}

func (e events) Less(i, j int) bool {
	return e[i].StartTime < e[j].StartTime
}

func (e events) Swap(i, j int) {
	buf := e[i]
	e[i] = e[j]
	e[j] = buf
}

type EventRepository struct {
	events events
}

func (e *EventRepository) AddEvent(newEvent domain.Event) error {
	for i, v := range e.events {
		if v.StartTime > newEvent.StartTime.Unix() {
			continue
		}

		if v.EndTime > newEvent.StartTime.Unix() {
			return errors.ErrDateBusy
		}

		if len(e.events)-1 > i && e.events[i+1].StartTime < newEvent.StartTime.Unix() {
			return errors.ErrDateBusy
		}

		e.events = append(e.events, Event{Id: newEvent.Id, Name: newEvent.Name,
			Description: newEvent.Description, StartTime: newEvent.StartTime.Unix(),
			EndTime: newEvent.EndTime.Unix()})

		sort.Sort(e.events)
		break

	}
	return nil
}

func (e *EventRepository) GetEvent(eventId int) (domain.Event, error) {
	for _, v := range e.events {
		if v.Id == eventId {
			return domain.Event{Id: v.Id, Name: v.Name, Description: v.Description,
				StartTime: time.Unix(v.StartTime, 0), EndTime: time.Unix(v.StartTime, 0)}, nil
		}
	}
	return domain.Event{}, errors.ErrEventNotFound
}

func (e *EventRepository) DeleteEvent(eventId int) error {
	for i, v := range e.events {
		if v.Id == eventId {
			copy(e.events[i:], e.events[:i+1])
			e.events = e.events[:len(e.events)-1]
			return nil
		}
	}
	sort.Sort(e.events)

	return errors.ErrEventNotFound
}

func (e *EventRepository) ChangeEvent(NewEvent domain.Event) error {
	for i, v := range e.events {
		if v.Id == NewEvent.Id {
			e.events[i] = Event{Id: NewEvent.Id, Name: NewEvent.Name,
				Description: NewEvent.Description, StartTime: NewEvent.StartTime.Unix(),
				EndTime: NewEvent.EndTime.Unix()}

			sort.Sort(e.events)

			return nil
		}
	}
	return errors.ErrEventNotFound
}

func (e *EventRepository) EventList() ([]domain.Event, error) {
	var events []domain.Event
	for _, v := range e.events {
		events = append(events, domain.Event{Id: v.Id, Name: v.Name,
			Description: v.Description, StartTime: time.Unix(v.StartTime, 0),
			EndTime: time.Unix(v.StartTime, 0)})
	}
	return events, nil
}
