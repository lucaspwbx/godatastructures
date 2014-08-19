package queue

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	if queue.Empty() != true {
		t.Errorf("Queue length should be 0")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	if q.Empty() != false {
		t.Errorf("Queue length should be 1")
	}
	if q.Front() != 1 {
		t.Errorf("Queue item should be 1")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Dequeue()
	if q.Empty() != true {
		t.Errorf("Queue length should be 0")
	}
}

func TestFront(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	if q.Front() != 1 {
		t.Errorf("Front item should be 1")
	}
}

func TestBack(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	if q.Back() != 2 {
		t.Errorf("Back item shoud be 2")
	}
}

func TestEmpty(t *testing.T) {
	emptyQueue, oneItemQueue := NewQueue(), NewQueue()
	oneItemQueue.Enqueue(1)
	var tests = []struct {
		q *Queue
		b bool
	}{
		{emptyQueue, true},
		{oneItemQueue, false},
	}

	for _, v := range tests {
		if v.q.Empty() != v.b {
			t.Errorf("Error")
		}
	}
}
