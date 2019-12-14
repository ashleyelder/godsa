package queue

import (
	"testing"
)

var q ItemQueue

func initQueue() *ItemQueue {
	q = ItemQueue{}
	q.New()
	return &q
}

func TestEnqueue(t *testing.T) {
	s := initQueue()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	// fmt.Println(reflect.TypeOf(s))
	// fmt.Println(reflect.TypeOf(s.slice))

	if itemsInQueue := len(s.slice); itemsInQueue != 3 {
		t.Errorf("wrong count, expected 3 but got %d", itemsInQueue)
	}
}

func TestDequeue(t *testing.T) {
	s := initQueue()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	s.Dequeue()
	if itemsInQueue := len(s.slice); itemsInQueue != 2 {
		t.Errorf("wrong count, expected 2 but got %d", itemsInQueue)
	}
	s.Dequeue()
	s.Dequeue()
	if itemsInQueue := len(s.slice); itemsInQueue != 0 {
		t.Errorf("wrong count, expected 0 but got %d", itemsInQueue)
	}
}
