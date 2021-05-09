package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}){
	//通过反射获取的传入变量的type,kind,值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=",rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=",n2)
	fmt.Printf("rval=%v,rval=%T\n",rVal,rVal)

	//下面我们将rval转成interface{}
	iv := rVal.Interface()
	//再把interface通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num",num2)
}

//结构体反射
func reflectTest02(b interface{}){
	//通过反射获取的传入变量的type,kind,值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=",rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	//3. 获取变量对应的kind
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind=%v kind=%v\n",kind1,kind2)

	//下面我们将rval转成interface{}
	iv := rVal.Interface()
	fmt.Printf("iV=%v, iV=%T\n",iv,iv)
	//再把interface通过断言转成需要的类型
	num2,ok := iv.(Student)
	if ok{
		fmt.Println("num",num2.Name)
	}
}

func reflect03(b interface{}){
	rVal := reflect.ValueOf(b)
	fmt.Printf("rval kind=%v\n",rVal.Kind())
	//rval
	rVal.Elem().SetInt(20)
	fmt.Println(rVal)
}

type Student struct{
	Name string
	Age int
}

func main(){
	//var num int = 100
	//reflectTest01(num)

	//定义student的实例
	stu := Student{
		Name:"vitas",
		Age:20,
	}
	reflectTest02(stu)

	const v int = iota
	fmt.Println(v)

	const(
		a=iota
		b = iota
		c,d = iota,iota
		e = iota
	)

	const(
		f=iota
		)
	fmt.Println(a,b,c,d,e,f)


	var num int = 10
	reflect03(&num)
	fmt.Println("num=",num)



}