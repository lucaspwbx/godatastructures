package stack

import "errors"

type Stack struct {
	dt []interface{}
}

func (s *Stack) isEmpty() bool {
	if len(s.dt) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(element interface{}) {
	s.dt = append(s.dt, element)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("stack is empty")
	}
	popped := s.dt[len(s.dt)-1]
	s.dt = s.dt[:len(s.dt)-1]
	return popped, nil
}

func (s *Stack) Peek() interface{} {
	if s.isEmpty() {
		return nil
	}
	return s.dt[len(s.dt)-1]
}

func (s *Stack) Size() int {
	return len(s.dt)
}

func NewStack() *Stack {
	return &Stack{}
}
