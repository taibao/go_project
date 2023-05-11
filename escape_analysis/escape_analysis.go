package escape_analysis

import "fmt"

//运行逃逸分析 go build -gcflags="-m -l" escape_analysis.go

func main() {
	//1 .函数类型逃逸分析
	//name := test()
	//fmt.Println(name())

	//2. interface{}  数据类型
	//name := "Golang"
	//fmt.Println(name) //println 方法传参类型interface，编译器对传入的变量类型未知，所有统一处理分配到了堆上

	//3. 指针 数据类型
	name := point()
	fmt.Println(*name)

}

// 函数类型逃逸分析
func test() func() string {
	return func() string {
		return "后端时光"
	}
}

// 指针 数据类型逃逸分析
func point() *string {
	name := "指针"
	return &name
}
