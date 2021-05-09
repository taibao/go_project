package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写程序
func test(key int){
	for i:=0;i<10;i++{
		fmt.Println("test() "+ strconv.Itoa(key) +" hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//在主线程（可以理解成进程）中，开启一个goroutine，该协程每隔一秒输出“hello world”
func main(){
	go test(1) //开启了一个协程
	go test(2) //开启了一个协程
	go test(3) //开启了一个协程
	go test(4) //开启了一个协程
	for i:=0;i<10;i++{
		fmt.Println("main() hello world " + strconv.Itoa(i))
	}
}