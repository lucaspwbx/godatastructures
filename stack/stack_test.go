package stack

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack.values != nil {
		t.Errorf("Stack values should be nil")
	}
}

func TestStackLength(t *testing.T) {
	stack := NewStack()
	if stack.length() != 0 {
		t.Errorf("Length should be 0")
	}
}

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stack.push(1)
	if stack.length() != 1 {
		t.Errorf("Stack should have length 1")
	}
	if stack.values[0] != 1 {
		t.Errorf("Stack item should be 1")
	}
}

func TestStackPop(t *testing.T) {
	stack := NewStack()
	stack.push(2)
	teste := stack.pop()
	if stack.length() != 0 {
		t.Errorf("Stack should have length 0")
	}
	if teste != 2 {
		t.Errorf("Item popped should be 2")
	}
}
