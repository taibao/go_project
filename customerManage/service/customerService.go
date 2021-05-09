package service

import "customerManage/model"

//该CustomerService，完成对Customer的操作，包括增删改查
type CustomerService struct{
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回*CustomerService,返回service对
func NewCustomerService() *CustomerService{
	//为了能够看到有客户在切片中
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"zhangsan","manewe",20000,"130668","wewefwe@1234.com")
	//将创建的客户存到客户切片中
	customerService.customers = append(customerService.customers,customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer{
	return this.customers
}

//添加客户到customer切片
//必须要用指针绑定*CustomerService，否则之前的service对象会丢失，导致内存泄露
func (this *CustomerService) Add(customer model.Customer) bool{
	//我们确定一个分配id的规则，添加的顺序
	this.customerNum++
	customer.Id  = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}

//根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int) bool{
	index := this.FindById(id)
	//如果index==-1，说明没有这个客户
	if index == -1{
		return false
	}

	//删除元素
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}

//修改数据
func (this *CustomerService) Edit(id int,name string,gender string,age int,phone string ,email string) bool{
	index := this.FindById(id)
	//如果index==-1，说明没有这个客户
	if index == -1{
		return false
	}
	return this.customers[index].Edit(name,gender,age,phone,email)
}

func (this *CustomerService) FindById(id int) int{
	index := -1
	//遍历this.customers 切片
	for i := 0;i< len(this.customers);i++{
		if this.customers[i].Id == id{
			//找到
			index = i
		}
	}

	return index
}






















