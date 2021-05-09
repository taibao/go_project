package main

import "fmt"

//如果结构体的字段类型是引用类型：指针，slice，map，零值都是nil，即没有分配空间
//如果需要使用这样的字段，需要先make，才能使用

//type Person struct{
//	Name string
//	Age int
//	Scores [5]float64
//	Ptr *int //指针 基本数据类型要new
//	Slice []int //切片
//	Map1 map[string]string //切片
//}


//type Cat struct{
//	Name string
//	Age int
//	Color string
//	hobby string
//}
//
//
//type A struct{
//	Num int
//}
//
//type B struct{
//	Num int
//}


func main(){
	//var a A
	//var b B
	//a = A(b) //若字段类型不一样，是不能强转的
	//
	//type C B
	//var c C
	//c = C(b)
	//fmt.Println(a,b,c)
	//
	//start := time.Now().Unix()
	//sum := 0
	//for i:=0;i<7000000000;i++{
	//	sum += i
	//}
	//end := time.Now().Unix()
	//
	//fmt.Println("  运行加法一共使用： ",sum,end-start)

	//创建一个结构体变量
	//var cat1  Cat
	//fmt.Printf("cat1的地址%p\n",&cat1)
	//cat1.Age = 19
	//cat1.Color = "red"
	//cat1.Name = "tom"
	//cat1.hobby = "吃<*)))><< "
	//fmt.Println(cat1)

	//var p1 Person
	////使用slice一定要先make
	//p1.Slice = make([]int,10)
	//p1.Slice[0] = 100
	////使用map一定要先make
	//p1.Map1 = make(map[string]string)
	//p1.Map1["name"] = "tom"
	//
	//var i int = 10
	//p1.Ptr = &i
	//fmt.Println(p1)

	//创建一个Monster变量

	//monster := Monster{"牛魔王",500,"开山斧"}
	//jsonstr,_ := json.Marshal(monster)
	//fmt.Printf("返回值：%v",string(jsonstr))

	//var vitas Person
	//vitas.Name = "nana"
	//vitas.speak()
	//vitas.count(10)
	//
	//res := vitas.count2(10)
	//
	//fmt.Println(res)

	//var c Circle
	//c.radius = 4.0
	//res := c.area()
	//fmt.Println("面积是=",res)

//指针传递方式
	var c Circle
	c.radius = 4.0
	//res2 := (&c).area2()
	//go语言编译器底层做了优化，(&c).area2()等价c.area2()
	//因为编译器会自动的加上&c
	res2 := c.area2()
	fmt.Printf("函数外c变量地址%p\n",&c)
	fmt.Println("radius = ",c.radius) //结构体内的字段已经发生了变化
	fmt.Println("面积是=",res2)

}


type Circle struct{
	radius float64
}

func (c Circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func (c *Circle) area2() float64{
	//因为c是指针，我们标准的访问字段的方式是(*c).radius
	//return 3.14 * (*c).radius * (*c).radius
	//(*c).radius等价c.radius
	fmt.Printf("函数内c变量地址%p\n",c) //与变量的地址完全一样
	c.radius = 10
	return 3.14 * c.radius * c.radius //编译器底层优化
}


















func (p Person) count2(n int) int{
	res := 0
	for i:=1;i<=n;i++{
		res += i
	}
	return res
}

func (p Person) count(n int){
	res := 0
	for i:=1;i<=n;i++{
		res += i
	}
	fmt.Println(p.Name,"计算的结果是",res)
}

func (p Person) speak(){
	fmt.Println(p.Name,"是个好人")
}

func (person Person) test(){
	person.Name = "vitas"
	fmt.Println("test() name = ", person.Name)
}

type Person struct{
	Name string
}



//type Monster struct{
//	Name string `json:"name"` //反射名字
//	Age int `json:"age"`
//	Weapon string `json:"weapon"`
//}