package main

import "fmt"

func main(){
	//管道可以声明为只读或者只写

	var chan2 chan<- int //声明为只写
	chan2 = make(chan int,3)
	chan2<-20
	fmt.Println("chan2=",chan2)


	//声明为只读
	var chan3 <- chan int
	num2 := <-chan3
	//chan3<-30 //错误
	fmt.Println("num2",num2)


}
