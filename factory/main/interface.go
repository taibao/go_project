package main

import (
	"fmt"
	"math/rand"
	"sort"
)
//
//type Usb interface {
//	Start()
//	Stop()
//}
//
//type Phone struct{
//}
//
////计算机
//type Computer struct{
//}
//
////编写一个方法working方法，接收一个usb接口类型变量
//func (c Computer) Working(usb Usb){
//	//只要实现了usb接口的所有方法，就能传入
//	usb.Start()
//	usb.Stop()
//}
//
////让phone实现usb接口的方法
//func (p Phone) Start(){
//	fmt.Println("手机开始工作.....")
//}
//func (p Phone) Stop(){
//	fmt.Println("手机停止工作....")
//}
//
//type Camera struct{
//
//}
////让camera实现Usb接口的方法
//func (c Camera) Start(){
//	fmt.Println("相机开始工作...")
//}
//func (c Camera) Stop(){
//	fmt.Println("相机停止工作....")
//}


type Stu struct{}

func (this *Stu) Start(){
	fmt.Println("start ...")
}

func (this *Stu) Stop(){
	fmt.Println("stop ...")
}



//让phone 实现USB接口的方法
func main(){
	//测试
	//先创建结构体变量
	//computer := Computer{}
	//phone := Phone{}
	//camera := Camera{}
	//
	//computer.Working(phone)
	//computer.Working(camera)
	//
	////空接口
	//var t interface{} = phone
	//var num1 float64 = 8.9
	//t = num1 //任何变量都可以赋给空接口
	//fmt.Println(t)

	//var stu Stu = Stu{}
	//var u Usb = &stu //不能用stu,因为是指针类型实现的
	//u.Start()
	//u.Stop()

	//定义一个数组切片
	var intSlice = []int{0,-1,10,7,90}
	sort.Ints(intSlice) //由于intSlice是引用类型，在函数内部排序后，变量值也改变了
	fmt.Println(intSlice)


	//查看实现的排序
	var heroes HeroSlice
	for i:=0;i<10;i++{
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d",rand.Intn(100)),
			Age: rand.Intn(100)+15,
		}
		heroes = append(heroes,hero)
	}

	fmt.Println(heroes)

	//调用sort的排序方法
	sort.Sort(heroes)
	fmt.Println(heroes)

}



//1：声明Hero结构体
type Hero struct{
	Name string
	Age int
}

//2：声明一个Hero结构体切片类型
type HeroSlice []Hero

//3：实现Interface接口
func (hs HeroSlice) Len() int{
	return len(hs)
}

//Less指定升序或降序排序
//1. 按Hero的年龄从小到大排序 !!
func (hs HeroSlice) Less(i,j int) bool{
	return hs[i].Age < hs[j].Age //按年龄排序
}

func (hs HeroSlice) Swap(i,j int){
	//temp := hs[i]
	//hs[i]  = hs[j]
	//hs[j] = temp
	//上面三句等价于
	hs[i],hs[j] = hs[j],hs[i]
}



