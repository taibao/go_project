package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体

type Monster struct{
	Name string `json:"name"`
	Age int
	Birthday string
	Sal float64
	Skill string
}

//结构体的序列化
func testStruct(){
	//演示
	monster := Monster{
		Name :"牛魔王",
		Age: 500,
		Birthday:"2021-01-10",
		Sal:9000,
		Skill:"法相天地",
	}
	//转为json
	data,err := json.Marshal(&monster)
	if err != nil{
		fmt.Printf("序列化失败,err=%v",err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v",string(data))
}

//将map进行序列化
func testMap(){
	//定义一个map
	var a map[string] interface{}
	//使用map，需要make
	a = make(map[string]interface{})
	a["name"]  = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	//将a这个map序列化
	data,err := json.Marshal(a) //map本身就是引用类型不用传&
	if err != nil{
		fmt.Printf("序列化失败,err=%v",err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v",string(data))
}

//对切片序列化，我们切片
func testSlice(){
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string] interface{})
	m1["name"]  = "红孩儿"
	m1["age"] = 30
	m1["address"] = "火云洞"
	slice = append(slice,m1)

	var m2 map[string]interface{}
	//使用map前，需要先make
	m2 = make(map[string] interface{})
	m2["name"]  = "白鹿公主"
	m2["age"] = 26
	m2["address"] = [2]string{"白鹿洞","火云洞"}
	slice = append(slice,m2)

	//切片序列化
	data,err := json.Marshal(slice) //slice本身就是引用类型不用传&
	if err != nil{
		fmt.Printf("序列化失败,err=%v",err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v",string(data))

}

func testFloat64(){
	var f float64 = 345.546 
	//切片序列化
	data,err := json.Marshal(&f)
	if err != nil{
		fmt.Printf("序列化失败,err=%v",err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v",string(data))
}

func main(){
	testStruct()
}
