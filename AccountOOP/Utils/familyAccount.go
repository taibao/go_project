package Utils

import (
	"fmt"
)

type FamilyAccount struct{
	//声明必要字段

	//声明一个变量,保存接收用户输入的选项
	key string

	//声明一个变量，控制是否退出for
	loop bool

	//定义账户的余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明`
	note string
	//定义一个变量，记录是否有收支的行为
	flag bool

	//收支的详情使用字符串来记录
	//当有收支时，只需对this.details进行拼接处理
	details string
}

//编写要给工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		key : "",
		loop : true,
		balance:10000.0,
		money:0.0,
		note:"",
		flag:false,
		details:"收支\t账户金额\t收支金额\t说	明",
	}
}

//显示明细写成一个方法
func (this *FamilyAccount) showdetails(){
	fmt.Println("--------------当前收支明细记录---------------")
	if this.flag{
		fmt.Println(this.details)
	}else{
		fmt.Println("当前没有收支明细....来一笔吧")
	}
}

//登记收入方法
func (this *FamilyAccount) income(){
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//将收入情况，拼接到this.details变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t3\t%v\t",this.balance,this.money,this.note)
	this.flag = true
}

//支出方法
func (this *FamilyAccount) pay(){
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance{
		fmt.Println("余额不足")

	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	//将收入情况，拼接到this.details变量
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t\t%v",this.balance,this.money,this.note)
	this.flag = true
}

//退出方法
func (this *FamilyAccount) exit(){
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for{
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n"{
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

//给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu(){
	for {
		fmt.Println("\n----------家庭收支记账软件---------")
		fmt.Println("          1：收支明细")
		fmt.Println("          2：登记收入")
		fmt.Println("          3：登记支出")
		fmt.Println("          4：退出软件")
		fmt.Println("----------请选择（1-4）：---------")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showdetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}
		if !this.loop {
			break;
		}

	}
}







