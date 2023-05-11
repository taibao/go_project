package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("23423")
	TestShops := strings.Split("appAKLWLitn7978,app38itOR341547", ",")
	fmt.Println("TestShops", TestShops)
	fmt.Println("message.AppId", "appAKLWLitn7978")
	if !InArray("appAKLWLitn7978", TestShops) {
		fmt.Println("未匹配店铺")
		return
	}
	fmt.Println("店铺匹配成功")
}

func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}
