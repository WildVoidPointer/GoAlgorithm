package stack

import "fmt"

func RunStackFunctionTest() {
	st := NewStack[int](10)
	if st == nil {
		return
	}

	st.Println()

	for i := 1; i <= st.size; i++ {
		st.Push(i)
	}
	st.Println()

	st.Push(233)

	for !st.IsEmpty() {
		popped := st.Pop()
		fmt.Printf(" Popped element: %v", popped)
	}

	st.Println()
}
