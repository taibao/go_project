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

func main() {
	strings := toSlice("hello", "world")
	fmt.Println("strings", strings)
	nums := toSlice(1, 2, 3, 4)
	fmt.Println("nums", nums)

}
