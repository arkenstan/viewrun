package ds

import "errors"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(a T) {
	s.data = append(s.data, a)
}

func (s *Stack[T]) Pop() error {
	if len(s.data) > 0 {
		s.data = s.data[:len(s.data)-1]
		return nil
	} else {
		return errors.New("Stack is empty")
	}
}

func (s *Stack[T]) Top() T {
	return s.data[len(s.data)-1]
}
