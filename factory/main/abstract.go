package main

import "fmt"

type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

//存钱
func (account *Account) Deposite(money float64,pwd string){
	//检验输入密码
	if pwd != account.Pwd{
		fmt.Println("密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0{
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

//取钱
func (account *Account) WithDraw(money float64,pwd string){
	//检验输入密码
	if pwd != account.Pwd{
		fmt.Println("密码不正确")
		return
	}

	//看看取款金额是否正确
	if money <= 0 || money > account.Balance{
		fmt.Println("你的余额不足")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

//查询余额
func (account *Account) Query(pwd string){
	//检验输入密码
	if pwd != account.Pwd{
		fmt.Println("密码不正确")
		return
	}
	fmt.Printf("%v你好,您的余额为：%v \n",account.AccountNo,account.Balance)
}



func main(){
	var acc Account
	acc.AccountNo="vitas"
	acc.Pwd= "123"
	acc.Deposite(300,"123")
	acc.Query("123")
	acc.WithDraw(150,"123")
}
