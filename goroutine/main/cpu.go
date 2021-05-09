package main

import (
	"fmt"
	"runtime"
)

func main(){

	//获取当前系统的cpu数量
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println("num=",num)
}

