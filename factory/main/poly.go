package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

//计算机
type Computer struct{
}
//编写一个方法working方法，接收一个usb接口类型变量
func (c Computer) Working(usb Usb){
	//只要实现了usb接口的所有方法，就能传入
	usb.Start()
	//若类型是手机还要调用call
	if phone,ok := usb.(Phone);ok{
		phone.Call()
	}
 	usb.Stop()
}


type Phone struct{
}

//让phone实现usb接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作.....")
}
func (p Phone) Stop(){
	fmt.Println("手机停止工作....")
}

func (p Phone) Call(){
	fmt.Println("手机来电话了....")
}
type Camera struct{

}
//让camera实现Usb接口的方法
func (c Camera) Start(){
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop(){
	fmt.Println("相机停止工作....")
}

//多态在go中主要用接口实现
func main(){
	var usbArr [3]Usb
	usbArr[0] = Phone{}
	usbArr[1] = Phone{}
	usbArr[2] = Camera{}

	//遍历usbArr，针对phone还有一个特有方法call，遍历usb数组
	//出了调用usb接口声明的方法外，还需要调用Phone特有方法call=>类型断言、


	var computer Computer
	for _,v:=range usbArr{
		computer.Working(v)
		fmt.Println()
	}

}