package model

import "fmt"

//声明customer结构体，表示客户信息

type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写工厂模式，返回一个customer实例
func NewCustomer(id int,name string,gender string,age int,phone string ,email string) Customer{
	return Customer{
		Id:id,
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}

//编写工厂模式，返回一个customer实例
func NewCustomer2(name string,gender string,age int,phone string ,email string) Customer{
	return Customer{
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}

//返回用户信息
func (this Customer) GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",this.Id,this.Name,this.Gender,
		this.Age,this.Phone,this.Email)
	return info
}

func (this *Customer) Edit(name string,gender string,age int,phone string ,email string) bool{
	this.Name = name
	this.Gender = gender
	this.Age  = age
	this.Phone = phone
	this.Email = email
	return true
}