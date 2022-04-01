package main

import (
	"fmt"
	"reflect"
)

type Animal struct{

}

func (n *Animal) Eat(){
	fmt.Println("EAt")
}

func (n *Animal) Dog(){
	fmt.Println("Dog")
}

func main(){
	animal := Animal{}
	value := reflect.ValueOf(&animal) //通过reflect反射获取对象信息

	f := value.MethodByName("Eat") //拿到反射的方法
	f.Call([]reflect.Value{}) //调用方法

	d := value.MethodByName("Dog")
	d.Call([]reflect.Value{})
}


