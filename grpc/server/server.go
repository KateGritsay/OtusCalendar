package server

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"time"
	"github.com/KateGritsay/OtusCalendar/event"
	"github.com/KateGritsay/OtusCalendar/pkg/grpc"
	)

type Server struct {
	calendar *event.Calendar
}

func NewServer(calendar *event.Calendar) *Server {
	return &Server{calendar}
}

func (server *Server) Create(_ context.Context, eventRequest *Event) (*Event, error) {
	date := time.Time{}
	if eventRequest.GetDate() != nil {
		tmp, err := ptypes.Timestamp(eventRequest.GetDate())
		if err != nil {
			return nil, err
		}
		date = tmp
	}

	var duration time.Duration
	if eventRequest.GetDuration() != nil {
		tmp, err := ptypes.Duration(eventRequest.GetDuration())
		if err != nil {
			return nil, err
		}
		duration = tmp
	}

	id := server.calendar.Add(event.Event{
		Date:        date,
		Duration:    duration,
		Description: eventRequest.GetDescription(),
	})
	eventRequest.Id = int64(id)

	return eventRequest, nil

}

