package stack

import "strconv"

type Stack struct {
	values []interface{}
}

func (s *Stack) push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) pop() interface{} {
	value := s.values[s.length()-1]
	s.values = s.values[:s.length()-1]
	return value
}

func (s *Stack) length() int {
	return len(s.values)
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) String() string {
	var a string
	for _, v := range s.values {
		if value, ok := v.(string); ok {
			a += value + "\n"
		}
		if number, ok := v.(int); ok {
			a += strconv.Itoa(number) + "\n"
		}
	}
	return a
}
