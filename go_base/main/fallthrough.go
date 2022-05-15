package main

import "fmt"

func main(){

	switch 1 {
	case 1:fallthrough
	case 2:
		fmt.Println("这是1或2") //使用了fallthrough可以直接跳过下一个case判断输出下一个结果
	case 4:
		fmt.Println("这是4")

	case 3: fmt.Println("这是3")

	}

}