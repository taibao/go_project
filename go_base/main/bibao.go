package main

import (
	"fmt"
	"strings"
)

var n int  = 20

//Name := "name" 等价于 var Name string ;Name = "tom" 赋值语句不能在函数外执行
var Name string = "name"


func main(){

	//f := makeSuffix(".jpg")
	//name := f("hello")
	//name2 := f("hello3")
	//
	//fmt.Println(name,name2)


	//res := sum(10,20)
	//fmt.Println("res =",res) //4.res = 30



	//test01()
	//test02()
	//test01()
fmt.Println(Name)

}


var name string = "tom"

func test01(){
	fmt.Println(name)
}

func test02(){
	name := "jack" //变量在函数内定义，不影响全局
	fmt.Println(name)
}


//defer
func sum(n1 int , n2 int) int {
	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=",n1) //defer 3: ok1 n1 = 10
	defer fmt.Println("ok2 n2=",n2) //defer 2 ok2 n2 = 20

	//增加一句话
	n1++ //11 但是defer的n1不变还是 10
	n2++ //21 但是defer的n2不变还是 20


	res := n1 + n2 //res = 30
	fmt.Println("ok3 res",res) // 1. ok3 res =30
	return res
}



//检查后缀
func makeSuffix(suffix string) func (string) string{
	return  func (name string) string{
		//如果name没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return name
	}
}