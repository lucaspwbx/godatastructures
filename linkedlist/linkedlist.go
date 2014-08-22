package linkedlist

type Node struct {
	element interface{}
	next    *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Find(item interface{}) *Node {
	currNode := l.head
	for {
		if currNode.element != item {
			if currNode.next == nil {
				return nil
				break
			}
			currNode = currNode.next
		} else {
			break
		}
	}
	return currNode
}

func (l *LinkedList) Insert(newElement interface{}, item interface{}) {
	newNode := NewNode(newElement)
	current := l.Find(item)
	newNode.next = current.next
	current.next = newNode
}

func (l *LinkedList) FindPrevious(item interface{}) *Node {
	currNode := l.head
	for {
		if currNode.next != nil && currNode.next.element != item {
			currNode = currNode.next
		} else {
			break
		}
	}
	return currNode
}

func (l *LinkedList) Remove(item interface{}) {
	prevNode := l.FindPrevious(item)
	if prevNode.next != nil {
		prevNode.next = prevNode.next.next
	}
}

func NewNode(element interface{}) *Node {
	return &Node{element, nil}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: NewNode("head")}
}
