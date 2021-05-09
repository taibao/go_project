package main

import (
	"customerManage/model"
	"customerManage/service"
	"fmt"
)

type customerView struct{
	key string //接收用户输入
	loop bool //表示是否循环的显示主菜单
	customerService *service.CustomerService //保存customerService
}

//显示所有客户信息
func (this *customerView) list(){
	//首先获取到当前所有客户信息
	customers :=this.customerService.List()
	//显示界面
	fmt.Println("---------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------客户列表完成---------------")
}

//得到用户的输入信息构建新的客户，并完成添加
func (this *customerView) add(){
	fmt.Println("------------添加客户-------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)

	//构建新的customer实例
	customer := model.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Add(customer) {
		fmt.Println("---------------添加完成------------------")
	}else {
		fmt.Println("---------------添加失败------------------")
	}
}

//得到用户输入的id，删除id对应的客户
func (this *customerView) delete(){
		fmt.Println("------------删除客户---------------")
		fmt.Println("请输入待删除客户编号（-1退出）")
		id := -1
		fmt.Scanln(&id)
		if id == -1 {
			return //放弃删除
		}
		fmt.Println("确认是否删除（y/n）：")

		//同学加入循环判断，知道用户输入y或n，才退出
		choice := ""
		fmt.Scanln(&choice)

		if choice == "y" ||choice == "Y"{
		//调用customerService的delete方法
		if this.customerService.Delete(id){
			fmt.Println("------------删除完成------------")
		}else{
			fmt.Println("------------删除失败，输入id不存在----------")
		}
	}
}

//退出软件
func (this *customerView) exit(){
	fmt.Println("确认是否退出（Y/N）：")
	for{
		fmt.Scanln(&this.key)
		if this.key =="Y" || this.key =="y" ||this.key == "N" || this.key == "n"{
			break
		}
		fmt.Println("你的输入有误，确认是否退出（Y/N）")
	}
	if this.key == "Y" || this.key == "y"{
		this.loop  = false
	}
}

func (this *customerView) edit(){
	//输入修改id
	fmt.Println("------------编辑客户---------------")
	fmt.Println("请输入待编辑客户编号（-1退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("你已退出编辑")
		return //放弃修改
	}
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)

	//构建新的customer实例
	if this.customerService.Edit(id,name,gender,age,phone,email) {
		fmt.Println("---------------修改完成------------------")
	}else {
		fmt.Println("---------------修改失败------------------")
	}
}

//显示主菜单
func (this *customerView) mainMenu(){
	for{
		fmt.Println("------------客户信息管理软件---------------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退    出")
		fmt.Println("            请选择（1-5）：")

		fmt.Scanln(&this.key)
		switch this.key{
		case "1":
			this.add()
		case "2":
			this.edit()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

		if !this.loop{
			break
		}
	}

	fmt.Println("你退出了客户关系管理系统")
}

func main(){
	//在main函数中，创建一个customerView,并运行显示主菜单
	customerView := customerView{
		key:"",
		loop:true,
	}
	//
	//显示主菜单..完成对customerView结构体customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()

}