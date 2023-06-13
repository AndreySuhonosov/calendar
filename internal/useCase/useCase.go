package useCase

import (
	"github.com/AndreySuhonosov/calendar/internal/domain"
	"github.com/AndreySuhonosov/calendar/internal/errors"
	"github.com/AndreySuhonosov/calendar/internal/interfaces"
)

type UseCase struct {
	repository interfaces.EventRepository
}

func (u *UseCase) AddEvent(event domain.Event) error {
	if _, err := u.repository.GetEvent(event.Id); err != errors.ErrEventNotFound {
		return err
	}

	err := u.repository.AddEvent(event)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) DeleteEvent(eventId int) error {
	err := u.repository.DeleteEvent(eventId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) ChangeEvent(NewEvent domain.Event) error {
	err := u.repository.ChangeEvent(NewEvent)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) EventList() ([]domain.Event, error) {
	list, err := u.repository.EventList()
	if err != nil {
		return nil, err
	}
	return list, nil
}
