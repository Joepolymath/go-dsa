package datastructures

import "fmt"

// stack using linkedlist

type Node struct {
	data	interface{}
	next	*Node
}

type Stack2 struct {
	head	*Node
	size	uint16
	MaxSize	uint16
}

func (s *Stack2) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack2) Push(value interface{}) (bool, error) {
	if s.size == s.MaxSize {
		return false, fmt.Errorf("Stack Overflow")
	}

	if s.head == nil {
		s.head = &Node{
			data: value,
		}
	} else {
		temp := s.head
		s.head = &Node{
			data: value,
			next: temp,
		}
	}

	s.size++
	return true, nil
}

func (s *Stack2) Pop() (interface{}, error) {
	if s.size == 0 {
		return nil, fmt.Errorf("Stack Underflow or Stack Empty")
	}

	temp := s.head
	s.head = temp.next
	s.size--
	return temp.data, nil
}

func (s *Stack2) Peek() (interface{}, error) {
	if s.size == 0 {
		return nil, fmt.Errorf("Stack Empty")
	}
	return s.head.data, nil
}

func (s *Stack2) GetSize() (uint16, error) {
	return s.size, nil
}

func (s *Stack2) GetBuffer() ([]interface{}, error) {
	buffer := make([]interface{}, s.size)
	temp := s.head
	for temp != nil {
		buffer = append(buffer, temp.data)
		temp = temp.next
	}
	return buffer, nil
}

