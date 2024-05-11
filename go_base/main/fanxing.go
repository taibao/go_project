package main

import (
	"fmt"
	"strings"
)

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

//泛型约束
//在使用泛型时，有时候需要对泛型类型进行一定的约束。例如，我们希望某个泛型函数或类型只能接受特定类型的参数，或者特定类型的参数必须实现某个接口，在go中，可以使用泛型约束来实现需求

// 类型榆树
// 类型榆树可以让泛型函数或类型只接受特定类型的参数，类型约束可以使用interface{}类型和类型断言来实现。
// 以下泛型函数，可以接受fmt.Stringer接口类型
//func Print[T fmt.Stringer](x T) {
//	fmt.Println(x.String())
//}

type StringWrapper string

func (s StringWrapper) String() string {
	return string(s)
}

// 约束语法
// 类型约束可以使用在类型参数后加上一个约束类型实现
//func Print[T fmt.Stringer, U io.Reader](x T, y U) {
//	fmt.Println(x.String())
//	_, _ = io.Copy(os.Stdout, y)
//}

type MyType[T fmt.Stringer] struct {
	data T
}

func (m *MyType[T]) String() string {
	return m.data.String()
}

// 该函数可交换任意两个变量的值
func Swap[T any](a, b, c *T) {
	*a, *b, *c = *c, *a, *b
}

//泛型接口
//泛型接口是一种可以处理多种类型数据的接口，在golang中，可以使用类型参数实现泛型接口

type Container[T any] interface {
	Len() int
	Add(T)
	Remove() T
}

// 泛型接口约束
// 泛型接口约束用于限制实现泛型接口的类型的范围，确保泛型代码只能用于满足特定条件的类型，在golang中，泛型接口约束使用接口来定义
type Stringer interface {
	String() string
}

// 接口被约束为只能存储实现了Stringer接口的类型
type ContainerV1[T Stringer] interface {
	Len() int
	Add(T)
	Remove() T
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
	//Print[StringWrapper]("hello world")
	a := 2
	b := 3
	c := 5

	Swap[int](&a, &b, &c)

	fmt.Println(a, b, c)

	//numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	//Sort(numbers)
	//fmt.Println(numbers)

	words := []string{"apple", "banana", "cherry"}
	upper := Map(words, func(word string) string {
		return strings.ToUpper(word)
	})
	fmt.Println(upper)

}

// 泛型使用场景
// 1排序
//func Sort[T comparable](s []T) {
//	sort.Slice(s, func(i, j int) bool {
//		return s[i] < s[j]
//	})
//}

// 搜索
//func Search[T comparable](s []T, x T) int {
//	return sort.Search(len(s), func(i int) bool {
//		return s[i] >= x
//	})
//}

// 映射
func Map[k comparable, V any](s []k, f func(k) V) map[k]V {
	result := make(map[k]V)
	for _, k := range s {
		result[k] = f(k)
	}
	return result
}
