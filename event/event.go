package event

import (
	"sync"
	"time"
)

type Event struct {
	Date        time.Time
	Duration    time.Duration
	Description string
}

type ID uint64

type Calendar struct {
	mtx    sync.RWMutex
	events map[ID]Event
	id     ID
}

func NewCalendar() *Calendar {
	return &Calendar{events: make(map[ID]Event), id: 1}
}

func (calendar *Calendar) CreateEvent(event Event) ID {
	calendar.mtx.Lock()
	id := calendar.id
	calendar.events[id] = event
	calendar.id++
	calendar.mtx.Unlock()
	return id
}

func (calendar *Calendar) UpdateEvent(id ID, event Event) (ok bool) {
	calendar.mtx.Lock()
	defer calendar.mtx.Unlock()
	_, ok = calendar.events[id]
	if !ok {
		return ok
	}
	calendar.events[id] = event
	return ok
}

func (calendar *Calendar) RemoveEvent(id ID) (ok bool) {
	calendar.mtx.Lock()
	_, ok = calendar.events[id]
	delete(calendar.events, id)
	calendar.mtx.Unlock()
	return ok
}

func (calendar *Calendar) GetEvent(id ID) (event Event, ok bool) {
	calendar.mtx.RLock()
	event, ok = calendar.events[id]
	calendar.mtx.RUnlock()
	return event, ok
}

