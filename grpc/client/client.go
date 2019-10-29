package main

import (
	"context"
	"fmt"
	calendarpb "github.com/KateGritsay/OtusCalendar/pkg/grpc"
	duration2 "github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"time"

)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
}

	c := calendarpb.NewCalendarClient(cc)
	id := createEvent(c)
	event := getEvent(c, id)
	fmt.Printf("Created event: %+v\n", event)

	event = updateEvent(c, event)

	fmt.Printf("Updated event description: %+v\n", event)

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()


func CreateEvent (c calendarpb.CalendarClient) uint64{
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	date := time.Now().Add(5 *time.Minute)
	duration := date.Add(5 *time.Minute)

	res, err := c.CreateEvent(ctx, &calendarpb.Event{
		Id:1,
		Date:&timestamp.Timestamp{Seconds: date.Unix(), Nanos: 0},
		Duration:&duration2.Duration{Seconds:duration.Unix(), Nanos:0},
		Description:"New description",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return res.GetId()

}

	func GetEvent (c calendarpb.CalendarClient, id  int64) *calendarpb.Event {
		ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
		defer cancel()

		res, err := c.GetEvent(ctx, &calendarpb.ID{Id: id})

		if err != nil {
			log.Fatal(err)
		}

		return res.GetEvent()
	}


	func RemoveEvent (c calendarpb.CalendarClient, id  int64) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	res, err := c.RemoveEvent(ctx, &calendarpb.ID{Id:id})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	return res.GetOk()
}

	func UpdateEvent (c calendarpb.CalendarClient, event *calendarpb.Event) bool{
		ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
		defer cancel()

		res, err := c.UpdateEvent(ctx, &calendarpb.Event{
			Id: 		 event.Id,
			Date:        event.Date,
			Duration:    event.Duration,
			Description: "Updated description",
			})
		if err != nil {
			log.Fatal(event)
		}

	return res.GetOk()
	}