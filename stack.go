package main

import "fmt"
import "strconv"

type Stack struct {
  structure []interface{}
}

func (s *Stack) push(value interface{}) {
  s.structure = append(s.structure, value)
}

func (s *Stack) pop() interface{} {
  s.structure = s.structure[:s.length()-1]
  return s.structure[s.length()-1]
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

func main() {
  stringStack := NewStack()
  stringStack.push("bla")
  stringStack.push("teste")

  fmt.Println(stringStack)

  intStack := NewStack()
  intStack.push(1)
  intStack.push(3)
  intStack.push(5)

  fmt.Println(intStack)

  intStack.pop()

  fmt.Println(intStack)
}
