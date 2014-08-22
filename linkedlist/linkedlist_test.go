package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	list := NewLinkedList()
	list.Insert("new", "head")
	node := list.Find("new")
	if node.element != "new" {
		t.Errorf("Error")
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkedList()
	list.Insert("new", "head")
	list.Remove("new")
	node := list.Find("new")
	if node != nil {
		t.Errorf("Error")
	}
}

func TestFindPrevious(t *testing.T) {
	list := NewLinkedList()
	list.Insert("new", "head")
	node := list.FindPrevious("new")
	if node.element != "head" {
		t.Errorf("Error")
	}
}
