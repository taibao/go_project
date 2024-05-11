package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"
)

// 时间和日期相关函数
func main() {
	//str := []byte("eyJjaGVja19hcHBfYXV0aCI6dHJ1ZSwib2JqIjpbXSwicmVnaW9uIjoiYXBwb2kyZ3lueWQ5ODg1Iiwic3ViIjoidV82Mjc2NDRmZjJkMThmX3FNeW0xUno2aTUifQ==")
	//fmt.Println("fsd", string(str))
	//test4()
	//GetCategoryTabSetting()
	test5()
}

func test5() {
	str := "https://appe0mes6qx8480.h5.xiaoeknow.com/v1/goods/order?id=SPU_ENT_1704961750YKKm3KoLeMjaO&spu_type=ENT&resource_id=g_659fa6d649d68_lrfXPLQw99&resource_type=21&sku=SKU_ENT_1704961750LP5fLiELfo&count=1"
	parsedURL, _ := url.Parse(str)
	fmt.Println(parsedURL)
	queryParams := parsedURL.Query()
	//resourceId := queryParams.Get("resource_id")
	//resourceType := queryParams.Get("resource_type")
	//Id := queryParams.Get("id")
	//spuType := queryParams.Get("spu_type")
	//sku := queryParams.Get("sku")
	//count := queryParams.Get("count")
	//
	fmt.Println("queryParams:", queryParams)
	//fmt.Println("resource_id:", param2)

}

func GetCategoryTabSetting() {

	cateGoryStr := `[{"tabTag":"home","subTabs":[{"title":"tab0","url":"https://appe0MEs6qX8480.h5.xiaoeknow.com"},{"title":"tab1","url":"https://appe0MEs6qX8480.h5.xiaoeknow.com/p/decorate/page/eyJpZCI6IjQ2NjA1MTUifQ"}]},{"tabTag":"vip","subTabs":[{"title":"tab0","url":"https://appe0MEs6qX8480.h5.xiaoeknow.com"},{"title":"tab1","url":"https://appe0MEs6qX8480.h5.xiaoeknow.com/p/decorate/page/eyJpZCI6IjQ2NjA1MTUifQ"}]}]`
	if cateGoryStr == "" {
		return
	}
	var arr []map[string]interface{}
	err := json.Unmarshal([]byte(cateGoryStr), &arr)
	if err != nil {
		return
	}
	fmt.Println(arr)
}
func test4() {
	data := []int{1, 2, 3, 4, 5, 6, 45, 7, 8, 9, 10}
	var sum int
	var wg sync.WaitGroup
	dataChan := make(chan int)
	defer close(dataChan)
	//获取
	for _, item := range data {
		wg.Add(1)
		item := item
		go func() {
			defer wg.Done()
			if item == 45 {
				dataChan <- 0
				return
			}
			dataChan <- item
		}()
	}

	for i := 1; i <= len(data); i++ {
		sum += <-dataChan
	}
	wg.Wait()

	fmt.Println("sum", sum)
}

func test3() {
	result := make(chan string)

	fmt.Println("同步成功")

	go func() {
		// 后台任务
		for i := 1; i <= 5; i++ {
			fmt.Println("执行任务", i)
		}

		// 提前返回OK
		result <- "OK"
	}()

	// 等待后台任务执行完毕或超时
	select {
	case res := <-result:
		return
		fmt.Println(res)
		//case <-time.After(10 * time.Second):
		//	fmt.Println("超时")
	}
}

func test2() {
	var wg sync.WaitGroup
	dataChan := make(chan []int)
	defer close(dataChan)

	page := 1
	size := 2
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var tem []int
	for {
		wg.Add(1)
		go func() {
			dataSlice := ArraySlice(data, uint((page-1)*size), uint(size))
			defer wg.Done()
			fmt.Println("dataSlice", dataSlice, "size", size)
			tem = dataSlice
			dataChan <- dataSlice
		}()
		time.Sleep(1 * time.Second)
		if len(tem) < size {
			fmt.Println(len(tem), size)
			break
		}
		page++
	}

	for i := 1; i <= page; i++ {
		resChunk := <-dataChan
		if resChunk == nil {
			continue
		}

		if len(resChunk) > 0 {
			//将查出的数据写入商品表
			fmt.Println("输出", resChunk)
		}
	}

	wg.Wait()
}

func test1() {
	data := []int{}
	if len(data) == 0 {
		return
	}

	var wg sync.WaitGroup
	dataChan := make(chan []int)
	defer close(dataChan)

	page := 1
	size := 3
	var courseAvailable []int
	for {
		if len(data) < size {
			size = len(data)
		}
		//批量查询权益
		dataSlice := ArraySlice(data, uint((page-1)*size), uint(size))
		fmt.Println("dataSlice", dataSlice, "size", size)
		wg.Add(1)
		go func() {
			defer wg.Done()
			courseAvailableSlice := GetDataSlice(dataSlice)
			dataChan <- courseAvailableSlice
		}()

		if page*size >= len(data) {
			break
		}
		page++
	}

	fmt.Println("page的值", page)

	for i := 1; i <= page; i++ {
		resChunk := <-dataChan
		if resChunk == nil {
			continue
		}

		if len(resChunk) > 0 {
			courseAvailable = append(courseAvailable, resChunk...)
		}
	}
	wg.Wait()

	fmt.Println(courseAvailable)

}

func GetDataSlice(data []int) []int {
	return data
}

func ArraySlice(s []int, offset, length uint) []int {
	if offset > uint(len(s)) {
		fmt.Println(offset, len(s))
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < uint(len(s)) {
		return s[offset:end]
	}
	return s[offset:]
}
