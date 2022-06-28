package main

import (
	"fmt"
)


func ArrayColumn(data []interface{},key string) []interface{}{
	var res []interface{}
	for _,item := range data{
		if item.(map[string]interface{})[key] == nil {
			panic("数组中不存在该key"+key)
		}
		appId := item.(map[string]interface{})[key].(string)
		res = append(res, appId)
	}
	return res
}

func main(){
	//var arr interface{}
	//

	var arr1 []interface{}
	arr1 = append(arr1,map[string]interface{}{
		"user_id":1,
		"updated_at" : "2022-02-17 19:39:35",
	})

	arr1 = append(arr1, map[string]interface{}{
		"user_id":2,
		"updated_at" : "2022-02-17 19:39:31",
	})

	arr1 = append(arr1,  map[string]interface{}{
		"user_id":3,
		"updated_at" : "2022-02-18 19:39:30",
	})

	updateAts := ArrayColumn(arr1,"updated_at")

	fmt.Println(updateAts)
}
