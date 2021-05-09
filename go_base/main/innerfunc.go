package main

import (
	"errors"
	"fmt"
)

func main(){
	//num1 :=100
	//fmt.Printf("num1的类型%T,num1的值%v，num1的地址%v",num1,num1,&num1)

	//num2 := new(int) //创建一个int指针
	//*num2 = 100
	//fmt.Printf("num2的类型%T,num2的值%v，num2的地址%v，指针指向的值=%v",num2,num2,&num2,*num2)


//test()
	testerr()
fmt.Println("程序继续执行")
}

//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确，就返回一个自定义的错误
func readConf(name string)(err error){
	if name == "config.ini"{
		return nil
	}else{
		return errors.New("读取文件错误..") //自定义错误信息
	}
}

func testerr(){
	err := readConf("config.ini")
	if err != nil{
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("testerr继续执行")
}

func test(){
	//使用defer+recover来捕获和处理异常
	defer func(){
		err := recover() //recover（）内置函数，可以捕获到异常
		if err != nil{ //说明捕获到错误
			fmt.Println("哎哟，报错误了，err=",err)
		}
	}() //调用匿名函数后面直接加()

	num1 := 10
	num2 := 0
	res := num1/num2
	fmt.Println(res)
}