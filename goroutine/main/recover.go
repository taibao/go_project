package main

import (
	"fmt"
	"time"
)

func sayHello(){
	for i := 0;i<10;i++{
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

//函数
func test3(){
	//使用recover捕捉错误
	defer func() {
		if err := recover();err != nil{
			fmt.Println("test()发生错误",err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0]  = "golang" //error
}

func main(){
	go sayHello()
	go test3()

	for i := 0;i<10;i++{
		fmt.Println("main() ok=",i)
		time.Sleep((time.Second))
	}
}
