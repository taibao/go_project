package main

import (
	"encoding/json"
	"fmt"
)

// 最内层
type lotsStatus struct { //车位信息
	ParkingSpaceno     string  `json:"parkingspaceno"`     //车位编号
	SpaceType          int     `json:"spacetype"`          //车位类型  1：普通车位
	ParkingSpaceStatus string  `json:"parkingspacestatus"` //车位状态 Y：占用（有车） N：未占用（无车）
	PlateNo            string  `json:"plateno"`            //占用车牌号
	Timestamp          float64 `json:"timestamp"`          //上传时间戳
}

// 包一层
type dataOne struct {
	DataOne []lotsStatus `json:"data"`
}

// 再包一层
type dataTwo struct {
	DataTwo dataOne `json:"data"`
}

// 主结构体，也就是最外面的一层
type LotsStatusAll struct {
	ReturnCode int     `json:"returncode"` //返回值
	Command    string  `json:"command"`    //命令
	Errordesc  string  `json:"errordesc"`  //是否成功
	Result     dataTwo `json:"result"`     //返回值
}


type phoneData struct{
	Id int `json:"id"`
	AppId string `json:"app_id"`
	ContentAppId string `json:"content_app_id"`
	ResourceId string `json:"resource_id"`
	UserId string `json:"user_id"`
	UniversalPhone string `json:"universal_phone"`
	State int `json:"state"`
	ResourceType int `json:"resource_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PhoneSubscribe struct{
	Data phoneData `json:"data"` //订阅值
	Op   string    `json:"op"`   //操作类型
}





func main() {
	var lotsStatus LotsStatusAll

	var strLotStatus = `{
	    "returncode": 0,
	    "command": "get all park state info",
	    "errordesc": "成功",
	    "result": {
	        "data": {
	            "data": [
	                {
	                    "parkingspaceno": "B34",
	                    "spacetype": 2,
	                    "parkingspacestatus": "1",
	                    "plateno": "",
	                    "timestamp": 1617296293.54021
	                },
	                {
	                    "parkingspaceno": "13Q-6",
	                    "spacetype": 2,
	                    "parkingspacestatus": "-1",
	                    "plateno": "",
	                    "timestamp": 1617296293.54021
	                }
	            ]
	        }
	    }
	}`
	err := json.Unmarshal([]byte(strLotStatus), &lotsStatus)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		//for _, v := range lotsStatus.Result.DataTwo.DataOne {
		//	fmt.Printf("ParkingSpaceStatus:%s  ParkingSpaceno:%s\n", v.ParkingSpaceStatus, v.ParkingSpaceno)
		//
		//}
	}

	str := `{
		"data":
			{
				"id":1125,
				"app_id":"apppcHqlTPT3482",
				"content_app_id":"",
				"resource_id":"i_5fed83f860b24642fed0b45f",
				"user_id":"u_5e562a8411cad_o6WwjIUsoP",
				"universal_phone":"13049353700",
				"state":1,
				"resource_type":1,
				"created_at":"2021-07-15 11:49:06",
				"updated_at":"1971-01-01 00:00:00"
			},
			"op":"+I"
		}`

	var phoneSubscribe PhoneSubscribe
	err = json.Unmarshal([]byte(str), &phoneSubscribe)
	if err != nil{
		fmt.Println("error:",err)
	}else {
		fmt.Println(phoneSubscribe)
	}


}