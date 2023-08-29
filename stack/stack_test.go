package stack

import "fmt"

func ExampleStack() {
	s := NewStack(1, 2)

	s.Push(3)

	fmt.Println(s.Count())
	fmt.Println(s.Peek())

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println(s.Count())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())

	// Output:
	// 3
	// 3 true
	// 3 true
	// 2 true
	// 1 true
	// 0
	// 0 false
	// 0 false
}
