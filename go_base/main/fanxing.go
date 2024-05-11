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

// 出栈
func (s *Stack[T]) Pop() T {
	n := len(s.data)
	x := s.data[n-1]
	s.data = s.data[:n-1]
	return x
}

func main() {
	//strings := toSlice("hello", "world")
	//fmt.Println("strings", strings)
	//nums := toSlice(1, 2, 3, 4)
	//fmt.Println("nums", nums)
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	x := s.Pop() //出栈
	fmt.Println(x)
}
