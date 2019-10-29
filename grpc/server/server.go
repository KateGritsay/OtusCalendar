package server

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
	"google.golang.org/grpc"
	"github.com/KateGritsay/OtusCalendar/event"
	calendarpb "github.com/KateGritsay/OtusCalendar/pkg/grpc"

	)

type Server struct {
	calendar *event.Calendar
}

func NewServer(calendar *event.Calendar) *Server {
	return &Server{calendar}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	calendarpb.RegisterCalendarServer(grpcServer, &Server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}

func (server *Server) CreateEvent(_ context.Context, eventReq *calendarpb.Event) (*calendarpb.Event, error) {
	date := time.Time{}
	if eventReq.GetDate() != nil {
		tmp, err := ptypes.Timestamp(eventReq.GetDate())
		if err != nil {
			return nil, err
		}
		date = tmp
	}

	var duration time.Duration
	if eventReq.GetDuration() != nil {
		tmp, err := ptypes.Duration(eventReq.GetDuration())
		if err != nil {
			return nil, err
		}
		duration = tmp
	}

	id := server.calendar.Add(event.Event{
		Date:        date,
		Duration:    duration,
		Description: eventReq.GetDescription(),
	})
	eventReq.Id = uint64(id)

	return eventReq, nil

}

func (server *Server) GetEvent(_ context.Context, id *calendarpb.ID) (*calendarpb.GetEventRes, error) {
	event, ok := server.calendar.Get(event.ID(id.Id))

	if !ok {
		return &calendarpb.GetEventRes{Result: &calendarpb.GetEventRes_Error{
			fmt.Sprintf("dont have event for id %d", id.Id),
		}}, nil
	}

	date, err := ptypes.TimestampProto(event.Date)
	if err != nil {
		return nil, err
	}
	duration := ptypes.DurationProto(event.Duration)

	return &calendarpb.GetEventRes{Result: &calendarpb.GetEventRes_Event{&calendarpb.Event{
		Id:          uint64(id.Id),
		Date:        date,
		Duration:    duration,
		Description: event.Description,
	}}}, nil
}
func (server *Server) RemoveEvent(_ context.Context, id *calendarpb.ID) (*calendarpb.UpdatedRes, error) {
	ok := server.calendar.Remove(event.ID(id.Id))

	if !ok {
		return &calendarpb.UpdatedRes{Result: &calendarpb.UpdatedRes_Error{
			fmt.Sprintf("dont have event for id %d", id.Id),
		}}, nil
	}

	return &calendarpb.UpdatedRes{Result: &calendarpb.UpdatedRes_Ok{ok}}, nil
}

func (server *Server) UpdateEvent(_ context.Context, eventReq *calendarpb.Event) (*calendarpb.UpdatedRes, error) {
	date := time.Time{}
	if eventReq.GetDate() != nil {
		tmp, err := ptypes.Timestamp(eventReq.GetDate())
		if err != nil {
			return nil, err
		}
		date = tmp
	}

	var duration time.Duration
	if eventReq.GetDuration() != nil {
		tmp, err := ptypes.Duration(eventReq.GetDuration())
		if err != nil {
			return nil, err
		}
		duration = tmp
	}

	id := event.ID(eventReq.GetId())

	ok := server.calendar.Update(id, event.Event{
		Date:        date,
		Duration:    duration,
		Description: eventReq.GetDescription(),
	})

	if !ok {
		return &calendarpb.UpdatedRes{Result: &calendarpb.UpdatedRes_Error{
			fmt.Sprintf("dont have event for id %d", id),
		}}, nil
	}

	return &calendarpb.UpdatedRes{Result: &calendarpb.UpdatedRes_Ok{ok}}, nil
}

