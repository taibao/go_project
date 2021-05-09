package main

import (
	"fmt"
)

//定义类型
type integer int

func (i integer) print(){
	fmt.Println("i=",i)
}

//编写一个方法，可以改变i的值
func (i *integer) change(){
	*i = *i + 1
}

type Student struct{
	Name string
	Age int
}

//Student 实现方法String()
func (stu *Student) String() string{
	str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}


type MethodUtils struct{
	//字段
}

//给MethodUtils
func (mu MethodUtils) Print(){
	for i := 1;i<= 10;i++{
		for j :=1;j <= 8; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//给MethodUtils
func (mu MethodUtils) Print2(m int,n int){
	for i := 1;i<= m;i++{
		for j :=1;j <= n; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
编写一个方法算该矩形的面积（可以接受len和宽width）
将其作为返回值，在main方法中调用该方法，接受返回面积值并打印
*/

func (mu MethodUtils) area(len float64 ,width float64) (float64){
	return len * width
}


//加减乘除

type Calcuator struct{
	Num1 float64
	Num2 float64
}


func (calcuator *Calcuator) getSum() float64{
	return calcuator.Num1 + calcuator.Num2
}

func (calcuator *Calcuator) getSub() float64{
	return calcuator.Num1 - calcuator.Num2
}

func (calcuator *Calcuator) getRes(operator byte) float64{
	res := 0.0
	switch operator{
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误...")
	}
	return res
}

func (calcuator *Calcuator) arr() {
	var arr = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	for i:=0;i<3;i++{
		fmt.Println(arr[i])
	}


	for i:=0;i<3;i++{
		for j:=i;j<3;j++{
			var tmp int
			tmp = arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = tmp
		}
	}
	for i:=0;i<3;i++{
		fmt.Println(arr[i])
	}

}

type Wife struct {
	Name string
}

func (p *Wife) test03(){
	p.Name = "nana"
	fmt.Println("test03() =",p.Name) //jack
}

func test05(p Wife){
	fmt.Println(p.Name)
}

//指针类型必须传地址
func test06(p *Wife){
	p.Name = "nana"
	fmt.Println(p.Name)
}


func main(){


	p := Wife{"vicky"}
	p.test03()
	fmt.Println("main() p.name=",p.Name)
	(&p).test03() //传地址也没法改变变量的值，编译器优化掉了
	fmt.Println("p的名字",p.Name)

	//指定结构体字段的三种方法
	var w Wife = Wife{"alice"}
	fmt.Println(w.Name)

	w2 := Wife{"alice"}
	fmt.Println(w2.Name)

	var w3 *Wife = &Wife{"lili"}
	var w4 *Wife = &Wife{Name:"hilay"}

	w3.test03()
	fmt.Println(w3.Name)
	fmt.Println(w4.Name)

	//方式2，返回结构体的指针类型
	//var stu2 = &Wife{"feifei"}



	//var cal Calcuator
	//cal.arr()
	//
	//
	//
	//cal.Num1 = 5
	//cal.Num2 = 4
	//res := cal.getRes('*')
	//fmt.Println(res)
	//
	//
	//
	//
	//var mu MethodUtils
	//mu.Print2(2,6)
	//
	//fmt.Println(mu.area(2.4,5.6))
	//Golang中的方法作用在指定的数据类型上，因此自定义类型都可以有方法
	//int，float32都可以有方法

	//var i integer = 10
	//i.change()
	//i.change()
	//i.change()
	//i.print()

	//定义student变量

	//stu := Student{
	//	Name:"tom",
	//	Age : 20,
	//}
	//
	//fmt.Println(&stu) //指针类型变量绑定了String，fmt.Println会默认调用
	////变量的String方法

}
