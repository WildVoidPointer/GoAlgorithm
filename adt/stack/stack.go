package stack

import "fmt"

type Stack[T any] struct {
	data   []T
	length int
	size   int
}

func NewStack[T any](size int) *Stack[T] {
	if size <= 0 {
		return nil
	}

	return &Stack[T]{
		data:   make([]T, 0, size),
		length: 0,
		size:   size,
	}
}

func (st *Stack[T]) IsEmpty() bool {
	return st.length <= 0
}

func (st *Stack[T]) IsFull() bool {
	return st.length >= st.size
}

func (st *Stack[T]) Peek() T {
	if st.IsEmpty() {
		var zero T
		return zero
	}
	return st.data[st.length - 1]
}

func (st *Stack[T]) Push(e T) bool {
	if st.IsFull() {
		return false
	}

	if st.length >= len(st.data) {
		st.data = append(st.data, e)
	} else {
		st.data[st.length] = e
	}

	st.length++

	return true
}

func (st *Stack[T]) Pop() T {
	if st.IsEmpty() {
		var zero T
		return zero
	}

	st.length--

	popped := st.data[st.length]
	var zero T
	st.data[st.length] = zero

	return popped
}

func (st *Stack[T]) Println() {
	fmt.Print("Stack {  ")
	for i := 0; i < st.length; i++ {
		fmt.Printf("%v  ", st.data[i])
	}
	fmt.Print("}\n")
}
