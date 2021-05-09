package main

import (
	"fmt"
	"strconv"
)

//演示golang
func main(){

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string //空的字符串

	//使用第一种方式来转换fmt. Sprintf方法
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q\n",str,str)

	//将char转为字符串

	str = fmt.Sprintf("%c",mychar) //char 转字符串
	fmt.Printf("str type %T str=%q\n",str , str)


	//第二种方式使用strconv 函数
	var num3 int =  99
	var num4 float64 = 23.456
	var b2 bool = true
	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str=%q\n",str , str)


	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str=%q\n",str , str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n",str , str)


}