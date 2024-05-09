package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//url := "https://fenxiao.xiaoe-tech.com/#/market/detail/appcctl2xam3820/p_658d9ba2e4b064a8fbdccde7/detail"
	//url := "https://app0cV8O7Pz9769.h5.xiaoeknow.com/p/course/column/p_62d14013e4b0eca59c165c53"
	//url := "https://appbgxmxhmv8623.h5.xiaoeknow.com/v1/goods/goods_detail/p_601bd138e4b05a9e88711c25?type=3"
	//url := "https://wzbiq.xetsl.com/s/2pzXLm"
	//if strings.Contains(url, "fenxiao.xiaoe-tech.com") {
	//	fmt.Printf("分销链接不能导入 '%s'\n", url)
	//	return
	//}
	//
	//if strings.Contains(url, "h5.xiaoeknow.com") {
	//	urlArr := strings.Split(url, "/")
	//	fmt.Println("链接切割", urlArr)
	//	domainArr := strings.Split(urlArr[2], ".")
	//	fmt.Println("domainArr", domainArr)
	//	AppId := domainArr[0]
	//	fmt.Println("app_id", AppId)
	//	resourceArr := strings.Split(urlArr[len(urlArr)-1], "?")
	//	ResourceId := resourceArr[0]
	//	fmt.Println("resource_id", ResourceId)
	//	return
	//}
	//fmt.Println("这是短链")

	intArr := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	zmArr := []interface{}{"a", "b", "c"}
	hzArr := []interface{}{"小", "宝"}
	ShuffleArr(intArr)
	fmt.Println(intArr)
	ShuffleArr(zmArr)
	fmt.Println(zmArr)
	ShuffleArr(hzArr)
	fmt.Println(hzArr)

	zmArr = append(zmArr, hzArr...)
	fmt.Println(zmArr)
	intArr = InsertList(intArr, zmArr, 2)
	fmt.Println(intArr)
}

func ShuffleArr(slice []interface{}) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func InsertList(course, objArr []interface{}, size int) []interface{} {
	var newRes []interface{}
	num := 0
	if size == 0 {
		size = 5
	}
	if course != nil && len(course) != 0 {
		key := 0
		for k, item := range course {
			if num >= size {
				num = 0
				key = k/size - 1
				if objArr != nil && len(objArr) != 0 {
					if key < len(objArr) {
						newRes = append(newRes, objArr[key])
					}
				}
			}
			newRes = append(newRes, item)
			num++
		}
		if objArr != nil && len(objArr) != 0 {
			for i := key + 1; i < len(objArr); i++ {
				newRes = append(newRes, objArr[i])
			}
		}
	} else {
		if objArr != nil && len(objArr) != 0 {
			newRes = append(newRes, objArr...)
		}
	}
	return newRes
}
