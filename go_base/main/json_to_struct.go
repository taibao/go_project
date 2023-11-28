package main

import (
	"encoding/json"
	"fmt"
)

type Alive struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Link int    `json:"link"`
}

func main() {
	var m interface{}
	str := `[{"name":"vitas","age":28}]`
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("解析失败")
		return
	}
	fmt.Println(m)

	var alive []Alive
	//var res *Alive
	res := MapToStruct(m, &alive)
	fmt.Println(*(res.(*[]Alive)))
}

func MapToStruct(m interface{}, alive interface{}) interface{} {
	mstr, err := json.Marshal(m)
	fmt.Println("map转字符串", string(mstr))
	err = json.Unmarshal(mstr, alive)
	if err != nil {
		fmt.Println("err2", err)
		return nil
	}
	fmt.Println("map转结构体", alive)
	return alive
}
