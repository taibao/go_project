package main

import "fmt"

type Point struct{
	x int
	y int
}

func main(){
	//var a interface{}
	//var point Point = Point{1,2}
	//a = point //ok
	////如何将a赋给一个Point变量
	//var b Point
	//b = a.(Point)  //类型断言 判断a是Point类型就赋值给b，不是就报错
	//fmt.Println(b)

	//类型断言的其他案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2

	//添加检测机制
	y,ok := x.(float64)
	if ok ==true{
		fmt.Println("类型转换成功，convert success")
	}else{
		fmt.Println("convert failed")
	}

	fmt.Println("程序执行完成",y)





}
