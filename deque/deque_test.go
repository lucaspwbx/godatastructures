package deque

import "testing"

func TestAddRear(t *testing.T) {
	expected := [2]int{3, 5}
	deque := NewDeque()
	deque.AddRear(5)
	deque.AddRear(3)
	if deque.Size() != 2 {
		t.Errorf("Size of deque should be 2")
	}
	for k, v := range deque.ds {
		if v != expected[k] {
			t.Errorf("Error")
		}
	}
}

func TestRemoveRear(t *testing.T) {
	deque := NewDeque()
	deque.AddRear(5)
	deque.AddRear(3)
	deque.RemoveRear()
	if deque.Size() != 1 {
		t.Errorf("Size of deque should be 1")
	}
	if deque.ds[0] != 5 {
		t.Errorf("Error")
	}
}

func TestAddFront(t *testing.T) {
	expected := [2]int{5, 3}
	deque := NewDeque()
	deque.AddFront(5)
	deque.AddFront(3)
	if deque.Size() != 2 {
		t.Errorf("Size of deque should be 2")
	}
	for k, v := range deque.ds {
		if v != expected[k] {
			t.Errorf("Error")
		}
	}
}

func TestRemoveFront(t *testing.T) {
	deque := NewDeque()
	deque.AddFront(5)
	deque.AddFront(3)
	deque.RemoveFront()
	if deque.Size() != 1 {
		t.Errorf("Size of deque should be 1")
	}
	if deque.ds[0] != 5 {
		t.Errorf("Error")
	}
}
