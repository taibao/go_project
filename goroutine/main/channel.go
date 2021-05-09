package main

import "fmt"

func main(){
	//演示管道的使用
	//创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)

	//2.看看intChan是啥
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n",intChan,&intChan)

	//向管道写入数据
	intChan <- 10
	num := 211
	intChan<-num

	fmt.Println(len(intChan),cap(intChan))

	var num2 int
	num2 = <-intChan //从管道取出一值
	num2 = <-intChan //从管道取出一值
	//num2 = <-intChan //从管道取出一值
	fmt.Println("num2",num2)
	fmt.Println(len(intChan),cap(intChan))

	//若数据全部取出，再取就会报deadlock


}
