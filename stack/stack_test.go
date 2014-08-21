package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	stack := NewStack()
	if stack.isEmpty() != true {
		t.Errorf("Should be empty")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	if stack.isEmpty() != false {
		t.Errorf("Should have one item")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	if item, _ := stack.Pop(); item != 1 {
		t.Errorf("Item should be 1")
	}
	if isEmpty := stack.isEmpty(); isEmpty != true {
		t.Errorf("Stack should be empty")
	}
	if _, err := stack.Pop(); err == nil {
		t.Errorf("should have returned error")
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	if item := stack.Peek(); item != 1 {
		t.Errorf("Peeked item should be 1")
	}
	if stack.isEmpty() != false {
		t.Errorf("Stack must remain the same after peeking of item")
	}
}

func TestSize(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	if size := stack.Size(); size != 1 {
		t.Errorf("Size shoud be 1")
	}
}
