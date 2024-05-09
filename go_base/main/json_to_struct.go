package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Alive struct {
	Name  string    `json:"name"`
	Age   int       `json:"age"`
	Link  int       `json:"link"`
	Birth time.Time `json:"birth"`
}

func main() {
	//mapToStruct()

	//structToMap()

	fmt.Println("Parsed time:", StrToTime("2023-07-19 10:00:00"))
}

func StrToTime(date string) time.Time {
	layout := "2006-01-02 15:04:05" // 布局必须与字符串的时间格式匹配
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Time{}
	}
	return t
}

func test12() {
	jsonData := `{"name":"John Doe","age":30,"birth":"1990-01-01 00:00:00"}`
	var person Alive
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(person)
}

// 结构体转map
func structToMap() {
	var m []Alive
	str := `[{"name":"vitas","age":28}]`
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("解析失败")
		return
	}
	fmt.Println(m)

	var alive []interface{}
	res := MapToStruct(m, &alive)
	fmt.Println(*(res.(*[]interface{})))
}

// map转结构体
func mapToStruct() {
	var m interface{}
	str := `[{"name":"vitas","age":28}]`
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("解析失败")
		return
	}
	fmt.Println(m)

	var alive []Alive
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
