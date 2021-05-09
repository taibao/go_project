package main

import (
	"fmt"
	"sort"
)

func main(){

	//var a map[string]string
	////在使用map前需要先make，make的作用就是给map分配数据空间
	//a = make(map[string]string,10) //可放10个key：value
	//a["no1"] = "宋江"
	//a["no2"] = "吴用"
	//a["no3"] = "卢俊义"
	//fmt.Println(a)
	//
	////第二种方式分配空间直接赋值
	//a = make(map[string]string)
	//a["no1"] = "宋江"
	//a["no2"] = "吴用"
	//fmt.Println(a)
	//
	////第三种方式
	//heroes := map[string]string{
	//	"hero1":"宋江",
	//	"hero2":"卢俊义"}
	//fmt.Println(heroes)


	//保存宿舍的学生姓名，性别
	//studentMap := make(map[string]map[string]string)
	//
	//studentMap["stu01"] = make(map[string]string,2)
	//studentMap["stu01"]["name"] = "tom"
	//studentMap["stu01"]["sex"] = "man"
	//
	//studentMap["stu02"] = make(map[string]string,2)
	//studentMap["stu02"]["name"] = "lili"
	//studentMap["stu02"]["sex"] = "girl"
	//fmt.Println(studentMap)

	////查找map
	//val , ok := studentMap["stu01"] //返回值和是否存在的布尔值
	//if ok{
	//	fmt.Printf("stu01存在,值为%v\n",val)
	//}else{
	//	fmt.Printf("没有stu01")
	//}
	//
	//delete(studentMap,"stu01")
	//fmt.Println(studentMap)
	//
	//
	////清空map
	//studentMap = make(map[string]map[string]string)
	//fmt.Println(studentMap)
	//

	////map的遍历
	//for k1,v1 := range studentMap{
	//	fmt.Println("k1=",k1)
	//	for k2,v2 := range v1{
	//			fmt.Printf("\t k2=%v v2=%v\n",k2,v2)
	//	}
	//	fmt.Println()
	//}
	//
	////统计map长度
	//fmt.Println(len(studentMap))

	//定义一个map切片

	//var monsters []map[string]string
	//monsters = make([]map[string]string,2) //准备放入两个妖怪
	//monsters[0] = make(map[string]string,2)
	//monsters[0]["name"] = "牛魔王"
	//monsters[0]["age"]	= "500"
	//
	//monsters[1] = make(map[string]string,2)
	//monsters[1]["name"] = "铁扇公主"
	//monsters[1]["age"]	= "300"
	//
	////这时已经没有小三的位置，再写会超出空间范围
	////想要再加得用append弹性扩容
	//newMonsters := make(map[string]string,2)
	//newMonsters["name"] = "玉面狐狸"
	//newMonsters["age"]	= "200"
	//monsters = append(monsters, newMonsters)
	//fmt.Println(monsters)


	map1 := make(map[int]int, 10)
	map1[10] =100
	map1[1] =13
	map1[4] =56
	map1[8] =90

	var keys []int
	for k,_:=range map1{
		keys = append(keys,k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _,k := range keys{
		fmt.Printf("map[%v]=%v",k,map1[k])
	}
//fmt.Println(map1)
}
