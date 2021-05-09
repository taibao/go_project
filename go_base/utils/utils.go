package utils

import "fmt"

//定义一个变量
var ImageText string = "图文"
var EleBook string = "电子书"
var GoodsBusiness = "商品服务"

func Echo(str string){
	fmt.Println(str)
}

func Test(n1 int,n2 int) (int,int){
	n1 = n1 + 1
	n2 = n2 + 1
	return n1,n2
}

func Test2(n int){
	if n > 0{
		n--
		Test2(n)
	}
	fmt.Println("n=",n)
}

func Fbn(n int) int {
	if(n ==1 || n==2){
		return 1;
	}else{
		return Fbn(n-1) + Fbn(n-2)
	}
}

func Hanshu(n int) int {
	if(n==1){
		return 3;
	}else{
		return 2*Hanshu(n-1)+1
	}
}

func Monkey(n int) int {
	if(n==1){
		return  1;
	}else{
		return (Monkey(n-1)+1)*2
	}
}

func Sum(n int) int {
	if(n==1){
		return 1
	}else{
		return n + Sum(n-1)
	}
}