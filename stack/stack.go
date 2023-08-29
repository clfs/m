// Package stack implements a generic stack.
package stack

// Stack is a generic stack implementation.
type Stack[T any] struct {
	data []T
}

// NewStack creates a new stack.
func NewStack[T any](data ...T) *Stack[T] {
	return &Stack[T]{data: data}
}

// Push pushes a value onto the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop pops a value from the stack.
func (s *Stack[T]) Pop() (t T, ok bool) {
	if len(s.data) == 0 {
		return
	}
	t = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	ok = true
	return
}

// Peek returns the top value from the stack without popping it.
func (s *Stack[T]) Peek() (t T, ok bool) {
	if len(s.data) == 0 {
		return
	}
	t = s.data[len(s.data)-1]
	ok = true
	return
}

// Count returns the number of elements in the stack.
func (s *Stack[T]) Count() int {
	return len(s.data)
}
