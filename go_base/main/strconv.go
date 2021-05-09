package main

import (
	"fmt"
	"strconv"
)

//演示golang
func main(){

	//var str string
	//var num5 int = 4567
	//str = strconv.Itoa(num5)
	//fmt.Printf("str type %T str=%q\n",str,str)

	var str string = "true"
	var b bool

	b, _ = strconv.ParseBool(str) //不想获取err，使用下划线忽略
	fmt.Printf("b type %T b = %v",b,b)

	var str2 string = "12345690"
	var n1 int64
	n1 ,_= strconv.ParseInt(str2,10,64)
	fmt.Printf("b type %T b = %v",n1,n1)
}