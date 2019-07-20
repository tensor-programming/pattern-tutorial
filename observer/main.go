package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	Event struct {
		data int
	}

	Observer interface {
		NotifyCallback(Event)
	}

	Subject interface {
		AddListener(Observer)
		RemoveListener(Observer)
		Notify(Event)
	}

	eventObserver struct {
		id   int
		time time.Time
	}

	eventSubject struct {
		observers sync.Map
	}
)

func (e *eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Observer: %d Recieved: %d after %v\n", e.id, event.data, time.Since(e.time))
}

func (s *eventSubject) AddListener(obs Observer) {
	s.observers.Store(obs, struct{}{})
}

func (s *eventSubject) RemoveListener(obs Observer) {
	s.observers.Delete(obs)
}

func (s *eventSubject) Notify(event Event) {
	s.observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		key.(Observer).NotifyCallback(event)
		return true
	})

}

func fib(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}

	}()

	return out
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

func main() {
	n := eventSubject{
		observers: sync.Map{},
	}

	t := time.Now()

	var obs1 = eventObserver{id: 1, time: t}
	var obs2 = eventObserver{id: 2, time: t}
	n.AddListener(&obs1)
	n.AddListener(&obs2)

	go func() {
		select {
		case <-time.After(time.Millisecond * 10):
			n.RemoveListener(&obs1)
		}
	}()

	for x := range fib(100000) {
		n.Notify(Event{data: x})
	}

}
