package main

import "fmt"

type Stack struct {
	data []string
}

// Push method : insert last item :

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

// is empty method :

func (s Stack) isEmpty() bool {
	return len(s.data) == 0
}

// get item by index :

func (s Stack) get(index int) string {
	if s.isEmpty() || index > len(s.data) || index < 0 {
		return ""
	}
	return s.data[index]
}

// pop methd : delete last item => FIFO :

func (s *Stack) Pop() string {
	if s.isEmpty() {
		return ""
	}
	index := len(s.data) - 1
	deleted_item := s.data[index]
	s.data = s.data[:index]
	return deleted_item
}

// display method :

func (s Stack) ShowStack() {
	itr := 0
	for itr != len(s.data) {
		fmt.Printf(" :-> %s ", s.data[itr])
		itr++
	}
}
