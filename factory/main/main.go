package main

import "fmt"

func main(){
	//var stu = model.Student{
	//	Name:"lina",
	//	Score:19.2,
	//}

	//var stu =  model.NewStudent("alice",19) //返回结构体指针stu
	//fmt.Println(*stu)
	////stu直接调用GetScore方法
	//fmt.Println("name=",stu.Name," score=",stu.GetScore())


	//nana := model.NewPerson("nana") //返回一个新分配的person指针
	//nana.SetAge(18)
	//nana.SetSal(12000)
	//
	//lucky := model.NewPerson("lucky")
	//lucky.SetAge(21)
	//lucky.SetSal(14000)
	//
	//fmt.Println(*nana)
	//fmt.Println(*lucky)

	//当我们对结构体嵌入了匿名结构体使用方法发生变化
	//p := &Pupil{}
	//p.Student.Name = "tom"
	//p.Student.Age = 10
	//p.testing()
	//p.Student.SetScore(80)
	//p.Student.ShowInfo()
	//
	//
	//g := &Graduate{}
	//g.Student.Name = "nana"
	//g.Student.Age = 20
	//g.testing()
	//g.Student.SetScore(80)
	//g.Student.ShowInfo()

	//var b B
	//b.Student.Name = "tom"
	//b.Age = 20
	//b.Score = 100
	//b.ShowInfo() //可简写
	//b.hello() //如果B自带Name就会打印B的Name，就算没有赋值，也会输出空字符串


 c := C{
 	&A{"tom","18"},
 	&B2{"nana","19"},
	 }
 fmt.Println(c.A.Name)
}

type C struct{
	*A
	*B2
}

//组合关系
type D struct{
	a A //有名结构体,字段不能直接访问
}

//编写学生考试系统
type Student struct {
	Name string
	Age int
	Score int
}

//将pupil和graduate共有的方法绑定到 *student
func (stu *Student) ShowInfo(){
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n",stu.Name,stu.Age,stu.Score)
}

func (stu *Student) SetScore(score int){
	stu.Score = score
}

func (stu *Student) GetScore() int{
	return stu.Score
}

type Graduate struct{
	Student //嵌入了Student匿名结构体
}
//特有方法
func (g *Graduate) testing(){
	fmt.Println("大学生正在考试")
}

type Pupil struct{
	Student
}
//特有方法
func (p *Pupil) testing(){
	fmt.Println("小学生正在考试")
}

//特有方法
func (p *Pupil) hello(){
	fmt.Println("P hello",p.Name)
}

type B struct {
	Pupil
	Name string
}

func (b *B) hello(){
	fmt.Println("b hello",b.Name)
}

type A struct {
	Name string
	Age string
}
type B2 struct {
	Name string
	Age string
}

