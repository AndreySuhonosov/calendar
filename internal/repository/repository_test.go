package repository

import (
	"github.com/AndreySuhonosov/calendar/internal/domain"
	"reflect"
	"testing"
	"time"
)

func TestEventRepository_AddEvent(t *testing.T) {
	type fields struct {
		events events
	}
	type args struct {
		newEvent domain.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "AddEvent", fields: fields{},
			args: args{newEvent: domain.Event{Id: 1, Name: "first", Description: "tests",
				StartTime: time.Now(), EndTime: time.Unix(time.Now().Unix()+1000, 0)}},
			wantErr: false,
		},
		{name: "AddExistEvent(startTime1 = startTime2)", fields: fields{events: events{Event{Id: 1, Name: "first", Description: "tests",
			StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 1000}}},
			args: args{newEvent: domain.Event{Id: 1, Name: "first", Description: "tests",
				StartTime: time.Now(), EndTime: time.Unix(time.Now().Unix()+1000, 0)}},
			wantErr: true,
		},
		{name: "AddExistEvent(EndTime1 > startTime2)", fields: fields{events: events{Event{Id: 1, Name: "first", Description: "tests",
			StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 1000}}},
			args: args{newEvent: domain.Event{Id: 1, Name: "first", Description: "tests",
				StartTime: time.Unix(time.Now().Unix()+200, 0), EndTime: time.Unix(time.Now().Unix()+1200, 0)}},
			wantErr: true,
		},
		{name: "AddExistEvent(EndTime1 < startTime2)", fields: fields{events: events{Event{Id: 1, Name: "first", Description: "tests",
			StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 1000}}},
			args: args{newEvent: domain.Event{Id: 1, Name: "first", Description: "tests",
				StartTime: time.Unix(time.Now().Unix()+1200, 0), EndTime: time.Unix(time.Now().Unix()+1300, 0)}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventRepository{
				events: tt.fields.events,
			}
			if err := e.AddEvent(tt.args.newEvent); (err != nil) != tt.wantErr {
				t.Errorf("AddEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_ChangeEvent(t *testing.T) {
	type fields struct {
		events events
	}
	type args struct {
		NewEvent domain.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{

		{name: "ChangeEvent", fields: fields{events: events{Event{Id: 1, Name: "first", Description: "tests",
			StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 1000}}},
			args: args{NewEvent: domain.Event{Id: 1, Name: "first", Description: "tests",
				StartTime: time.Now(), EndTime: time.Unix(time.Now().Unix()+1000, 0)}},
			wantErr: false,
		},
		{name: "ChangeNotFoundEvent", fields: fields{},
			args: args{NewEvent: domain.Event{Id: 1, Name: "first", Description: "tests",
				StartTime: time.Now(), EndTime: time.Unix(time.Now().Unix()+1000, 0)}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventRepository{
				events: tt.fields.events,
			}
			if err := e.ChangeEvent(tt.args.NewEvent); (err != nil) != tt.wantErr {
				t.Errorf("ChangeEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_DeleteEvent(t *testing.T) {
	type fields struct {
		events events
	}
	type args struct {
		eventId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "DeleteEvent", fields: fields{events: events{Event{Id: 1, Name: "first", Description: "tests",
			StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 1000}}},
			args:    args{eventId: 1},
			wantErr: false,
		},
		{name: "DeleteNotFoundEvent", fields: fields{events: events{Event{Id: 2, Name: "first", Description: "tests",
			StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 1000}}},
			args:    args{eventId: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventRepository{
				events: tt.fields.events,
			}
			if err := e.DeleteEvent(tt.args.eventId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_EventList(t *testing.T) {
	type fields struct {
		events events
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Event
		wantErr bool
	}{
		{name: "DeleteEvent", fields: fields{events: events{Event{Id: 1, Name: "first", Description: "tests",
			StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 1000}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventRepository{
				events: tt.fields.events,
			}
			got, err := e.EventList()
			if (err != nil) != tt.wantErr {
				t.Errorf("EventList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventRepository_GetEvent(t *testing.T) {
	type fields struct {
		events events
	}
	type args struct {
		eventId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventRepository{
				events: tt.fields.events,
			}
			got, err := e.GetEvent(tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_events_Len(t *testing.T) {
	tests := []struct {
		name string
		e    events
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_events_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		e    events
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_events_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		e    events
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Swap(tt.args.i, tt.args.j)
		})
	}
}
