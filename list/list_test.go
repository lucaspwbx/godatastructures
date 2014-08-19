package list

import "testing"

func TestInsert(t *testing.T) {
	l := NewList()
	l.Insert(1)
	if len(l.ds) != 1 {
		t.Errorf("Error inserting number")
	}
}

func TestFindPos(t *testing.T) {
	l := NewList()
	l.Insert(1)
	l.Insert(2)
	if index, _ := l.FindPos(2); index != 1 {
		t.Errorf("Error finding element 2")
	}
}

func TestRemove(t *testing.T) {
	l := NewList()
	l.Insert(1)
	l.Remove(1)
	if len(l.ds) != 0 {
		t.Errorf("Error removing element")
	}
}
