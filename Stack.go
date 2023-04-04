package main

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	last := len(s.data) - 1
	val := s.data[last]
	s.data = s.data[:last]
	return val
}

func (s *Stack) Peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Len() int {
	return len(s.data)
}
