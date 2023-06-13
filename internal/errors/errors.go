package errors

type EventError string

func (e EventError) Error() string {
	return string(e)
}

var (
	ErrDateBusy          = EventError("this date is busy")
	ErrIncorrectEndDate  = EventError("this date is incorrect")
	ErrEventAlreadyExist = EventError("this event already exist")
	ErrEventNotFound     = EventError("this event not exist")
)
