package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main(){
	content, err := os.ReadFile("export.json")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(content),"\n")

	appIds, err := os.ReadFile("app_ids.json")
	appIdArr := strings.Split(string(appIds),"\n")

	//var existApp []interface{}
	var existBUserId []interface{}

	for _,item := range arr{
		res := make(map[string]interface{})
		err = json.Unmarshal([]byte(item),&res)
		resContent := make(map[string]interface{})
		if Empty(res["content"]){
			continue
		}
		params := strings.Replace(res["params"].(string),"\\","",-1)
		resStr := strings.Replace(res["content"].(string),"\\","",-1)

		err = json.Unmarshal([]byte(resStr),&resContent)

		if !Empty(resContent["data"]) {
			var paramsData map[string]interface{}
			_ = json.Unmarshal([]byte(params),&paramsData)

			item2 := resContent["data"].([]interface{})
			for _,v :=range item2{
				if !Empty(v.(map[string]interface{})["app_id"]) && InArray(v.(map[string]interface{})["app_id"], appIdArr){
					//if !InArray(v.(map[string]interface{})["app_id"],existApp){
					//	fmt.Println(v.(map[string]interface{})["app_id"])
					//	existApp = append(existApp,v.(map[string]interface{})["app_id"])
					//}

					if !InArray(paramsData["b_user_id"],existBUserId){
						fmt.Println(paramsData["b_user_id"])
						existBUserId = append(existBUserId,paramsData["b_user_id"])
					}
				}
			}
		}
	}
}


// Empty empty()
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}


// InArray in_array()
// haystack supported types: slice, array or map
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
