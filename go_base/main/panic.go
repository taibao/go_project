package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)


func main(){
		//go func(){
		//	fmt.Println("hello")
		//	panic("这是例子")
		//}()

	//Go(func(){
	//	fmt.Println("hello")
	//	panic("这是例子")
	//})

	fmt.Println("vim-go")

	Go(func(){
		fmt.Println("hello")
		panic("这是例子")
	})

	time.Sleep(5*time.Second)




		//time.Sleep(5*time.Second)
	//err := 1/0
	//fmt.Errorf("输出错误 %v：%v,%w", "类型","0不能作为除数",err)

}





	func Go(x func()){
		go func(){
			defer func(){
				if err := recover(); err != nil{
					fmt.Println(err)
				}
			}()
			x()
		}()
	}
