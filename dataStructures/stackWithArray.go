package datastructures

import "fmt"

type Stack interface {
	Push(value interface{}) (bool, error)
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	GetBuffer() ([]interface{}, error)
	GetSize() (int32, error)
	IsEmpty() (bool, error)
}

type Stack1 struct {
	Size    int32
	Data    []interface{}
	MaxSize int32
}

func (s Stack1) Push(value interface{}) (bool, error) {
	if s.Size == s.MaxSize {
		return false, fmt.Errorf("Stack Overflow")
	}
	s.Data = append(s.Data, value)
	s.Size++
	return true, nil
}

func (s Stack1) Pop() (interface{}, error) {
	if s.Size == 0 {
		return nil, fmt.Errorf("Stack Underflow")
	}
	poppedElement := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return poppedElement, nil
}

func (s Stack1) Peek() (interface{}, error) {
	if s.Size == 0 {
		return nil, fmt.Errorf("Stack is empty")
	}
	lastElement := s.Data[len(s.Data) - 1]
	return lastElement, nil
}