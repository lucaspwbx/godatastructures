package list

import "testing"

func TestInsert(t *testing.T) {
	var data = []struct {
		item     interface{}
		expected int
	}{
		{1, 1},
		{"item", 1},
		{1.29, 1},
	}

	l := NewList()
	for _, v := range data {
		l.Insert(v)
		if len(l.ds) != v.expected {
			t.Errorf("Error inserting %v", v)
		}
		l.Clear()
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
	if l.Size() != 0 {
		t.Errorf("Error removing element")
	}
}

func TestClear(t *testing.T) {
	l := NewList()
	l.Insert(1)
	l.Clear()
	if l.Size() != 0 {
		t.Errorf("Error clearing list")
	}
}

func TestSize(t *testing.T) {
	l := NewList()
	if l.Size() != 0 {
		t.Errorf("Error, size should be 0")
	}
}
