package main

import (
	"fmt"
)

// Define a Stack struct type

type Stack struct {
	data []string
	length int
}

func (s *Stack) push(str string) {
	s.data = append(s.data, str)
	s.length += 1
}

func (s *Stack) pop() string {
	if s.length > 0 {
		element := s.data[s.length - 1]
		s.data = s.data[:s.length - 1]
		s.length = s.length - 1
	
		return element
	} else {
		return "STACK IS EMPTY"
	}
}

func (s *Stack) read() string {
	if s.length > 0 {
		return s.data[s.length - 1]
	}
	
	return "STACK IS EMPTY"
}

// The reverse function uses a stack to reverse the input string
func reverse(str string) (reversedStr string) {
	stack := &Stack{
		data: []string{},
		length: 0,
	}

	for _, runeVal := range str {
		stack.push(string(runeVal))
	}

	for i := stack.length; i > 0; i-- {
		reversedStr += stack.pop()
	}

	return
}

func main() {
	str := "abcde"
	fmt.Println(reverse(str))
}