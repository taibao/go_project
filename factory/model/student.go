package model

//定义一个结构体

//type Student struct{
//	Name string
//	Score float64
//}

type student struct{
	Name string
	score float64
}
//将studnet首字母改小写后,使用工厂模式
//返回student指针
func NewStudent(n string ,s float64) *student{
	return &student{
		Name:n,
		score:s,
	}
}

//如果score字段首字母小写，在其他包不可以直接访问，我们可以提供一个方法
func (s *student) GetScore() float64{
	return s.score
}
