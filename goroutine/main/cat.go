package main

import "fmt"

type Cat struct{
	Name string
	Age int
}

func main(){
	//定义一个存放任意数据类型的管道3个数据
	//var allChan chan interface{}

	allChan := make(chan interface{},3)
	allChan<-10
	allChan<-"tom jack"
	cat := Cat{"小花猫",4}
	allChan<-cat

	//获取第三个元素
	<-allChan
	<-allChan
	newCat := <-allChan
	tom := newCat.(Cat) //类型断言，将空接口转为相关结构体类型
	fmt.Println(tom.Name)
}