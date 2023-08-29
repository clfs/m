package stack

import "fmt"

func ExampleStack() {
	s := NewStack(7, 8)

	s.Push(9)

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
	// 9 true
	// 9 true
	// 8 true
	// 7 true
	// 0
	// 0 false
	// 0 false
}
