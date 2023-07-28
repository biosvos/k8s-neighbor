package lib

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

type Stack[T any] struct {
	elements []T
}

// Push change state
func (s *Stack[T]) Push(elements ...T) {
	s.elements = append(s.elements, elements...)
}

// Drop change state
func (s *Stack[T]) Drop(count int) {
	s.elements = s.elements[:len(s.elements)-count]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() uint64 {
	return uint64(len(s.elements))
}

func (s *Stack[T]) Peek() T {
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) All() []T {
	return s.elements
}
