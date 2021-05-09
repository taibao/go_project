package main

import "fmt"

func main(){
	intChan := make(chan int,3)
	intChan<-100
	intChan<-200
	close(intChan)

	//这是不能够再写入数到chanel
	//intChan<-300 //panic: send on closed channel
	fmt.Println("ok")
	//管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=",n1)

	//遍历管道
	intChan2 := make(chan int ,100)
	for i:=0;i<100;i++{
		intChan2<- i*2 //放100个数据到管道
	}
	close(intChan2) //关闭管道后就能正常遍历退出
	//遍历管道必须用foreach，因为管道长度实时变化
	for v:= range intChan2{
		fmt.Println("v=",v)
	}



}