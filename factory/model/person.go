package model

import "fmt"

type person struct{
	Name string
	age int
	sal float64
}

//工厂模式函数，返回实例
func NewPerson(name string) *person{
	return &person{
		Name:name,
	}
}

//为了访问age和sal，我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int){
	if age >0 && age <150{
		p.age = age
	}else{
		fmt.Println("年龄范围不正确")
		//给程序员一个默认值
	}
}

func (p *person) GetAge() int{
	return p.age
}

func (p *person) SetSal(sal float64){
		p.sal = sal
}

func (p *person) GetSal() float64{
	return p.sal
}
