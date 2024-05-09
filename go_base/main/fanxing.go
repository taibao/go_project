package main

import "fmt"

// 泛型函数
func toSlice[T any](args ...T) []T {
	return args
}

// 泛型类型
//type TypeName[T Type] struct {
//	//
//}

type Stack[T any] struct {
	data []T
}

// 入栈
func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

//出栈

func main() {
	strings := toSlice("hello", "world")
	fmt.Println("strings", strings)
	nums := toSlice(1, 2, 3, 4)
	fmt.Println("nums", nums)

}
