package main

import (
	"fmt"
)

func main(){
	//先输入数据个数
	var num int
	for{
		_,err := fmt.Scanf("0x%x",&num)
		if err !=nil{
			return
		}
		fmt.Println(num)
		break
	}
}
