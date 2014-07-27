package stack

import "strconv"

type Stack struct {
	structure []interface{}
}

func (s *Stack) push(value interface{}) {
	s.structure = append(s.structure, value)
}

func (s *Stack) pop() interface{} {
	value := s.structure[s.length()-1]
	s.structure = s.structure[:s.length()-1]
	return value
}

func (s *Stack) length() int {
	return len(s.structure)
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) String() string {
	var a string
	for _, v := range s.structure {
		if value, ok := v.(string); ok {
			a += value + "\n"
		}
		if number, ok := v.(int); ok {
			a += strconv.Itoa(number) + "\n"
		}
	}
	return a
}
