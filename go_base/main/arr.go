package main

import (
	"fmt"
	"strings"
	"time"
)

func GetMinutesAgo() string{
	currentTime := time.Now()
	t, _ := time.ParseDuration("-8h")
	result := currentTime.Add(t)
	return result.Format("2006-01-02 15:04:05.000000000")
}



func HttpBuildQuery(params map[string]string) (param_str string) {
	params_arr := make([]string, 0, len(params))
	for k, v := range params {
		params_arr = append(params_arr, fmt.Sprintf("%s=%s", k, v))
	}
	//fmt.Println(params_arr)
	param_str = strings.Join(params_arr, "&")
	return param_str
}


func main() {

	fmt.Println(GetMinutesAgo())

	//str := map[string][]string{"first": {"value"}, "multi": {"foo bar", "baz"}}
	//decodeStr = HttpBuildQuery(str)

	//str :="eyJhcHBfaWQiOiJhcHBvdnlqdm5leTI2MDEiLCJkYXRhIjp7Ind4X3VuaW9uX2lkIjoib1RIVzV2NmFnZzFiOGh4M1Z5YkM5V21VM3dhOCIsInd4X25pY2tuYW1lIjoi5Y2T5ZiJ5a6+Iiwid3hfYXZhdGFyIjoiaHR0cHM6Ly90aGlyZHd4LnFsb2dvLmNuL21tb3Blbi92aV8zMi9RMGo0VHdHVGZUSlJ2RFZUekRMNlZBaWEyRUtFd09Xd1o1NDJ5STM2QTdjb2hGbVBrYXJlRXY2Z01IOHZnOTl5aWFGdHNRM2liemZqMnd5OVpEUUxMRWJBQS8xMzIifX0="
	//fmt.Println(string([]byte(str)))
	//arr := map[string]interface{}{
	//	"app_id":"",
	//
	//}
	//
	//var uri url.URL
	//q := uri.Query()
	//q.Add("name", "张三")
	//q.Add("age", "20")
	//q.Add("sex", "1")
	//q.Add("wew","1232")
	//queryStr := q.Encode()
	//fmt.Println(queryStr)

	//fmt.Println(91/10)

	//测试break
	//for i := 0; i < 2; i++ {
	//	j := 0
	//	for {
	//		if j >= 5 {
	//			break
	//		}
	//		j++
	//		fmt.Println("第", i, "列")
	//	}
	//}


	//jsons := `["1：修复直播黑屏，app闪退","2：优化ui"]`
	//var msg  []string
	//err := json.Unmarshal([]byte(jsons), &msg)
	//
	//fmt.Println(msg , err)
	//
	//
	//
	//os.Exit(1)
	//使用数组求平均
	//
	//var hens [6]float64 //定义数组
	//
	//hens[0] = 3.0
	//hens[1] = 5.0
	//hens[2] = 1.0
	//hens[3] = 3.4
	//hens[4] = 2.0
	//hens[5] = 50.0
	//
	//totalWeight2 := 0.0
	//for i:=0;i<len(hens);i++{
	//	totalWeight2 += hens[i]
	//}
	//
	////平均体重
	//avgWeight2 := fmt.Sprintf("%.2f",totalWeight2/float64(len(hens)))
	//fmt.Printf("totalWeight2 %v ,avgWeight2 = %v,",totalWeight2,avgWeight2)

	//var intArr [3]int
	////当我们定义完数组后，其实数组的各个元素有默认值0
	//fmt.Println(intArr)
	//fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p",&intArr,&intArr[0],&intArr[1],&intArr[2])
	//

	//1:数组的地址可以通过数组名来获取&intArr
	//2.

	//var score [5]float64
	//
	//for i :=0;i<len(score);i++{
	//	fmt.Printf("请输入第%d个元素的值\n",i+1)
	//	fmt.Scanln(&score[i])
	//}
	//
	////变量数组打印
	//for i:=0;i<len(score);i++{
	//	fmt.Printf("score[%d]=%v",i,score[i])
	//}

	////四种初始化数组的方式
	//var numArr01 [3]int = [3]int{1,2,3}
	//fmt.Println("numArr01=",numArr01)
	//
	//var numArr02 = [3]int{5,6,7}
	//fmt.Println("numArr02=",numArr02)
	//
	//var numArr03 = [...]int{8,9,10} //三个点是固定写法，不能增减改变
	//fmt.Println("numArr03=",numArr03)
	//
	//var numArr04 = [...]int{1:800,0:900,2:299} //指定下标赋值
	//fmt.Println("numArr04=",numArr04)
	//
	////也可以类型推导
	//numArr05 :=  [...]int{1:800,0:900,2:299}
	//fmt.Println("numArr05=",numArr05)

	//heroes := [...]string{"宋江","吴用","卢俊义"}
	//
	//heroes2 := heroes
	//heroes2[0] = "卓嘉宾"
	//for i,v := range heroes2{
	//	fmt.Printf("heroes2[%d]=%v\n",i,v)
	//}
	//
	//for i,v := range heroes{
	//	fmt.Printf("heroes[%d]=%v\n",i,v)
	//}

	//var b [26]byte
	//index := 0
	//for i :='A';i<='Z';i++{
	//	b[index] = byte(i)
	//	index++
	//}
	//
	//fmt.Printf("输出值%c",b)

	//数组平均值
	//	var intArr = [...]int{1,-1,9,90,11}
	//	sum := 0
	//	for _,v := range intArr{
	//		//累计求和
	//		sum += v
	//	}
	//
	//	avg := float64(sum) /float64(len(intArr))
	//	fmt.Printf("sum=%v abg=%v",sum,avg)

	//var intArr3 [5]int
	////为了每次生成的随机数不一样，我们需要给一个seed值
	//rand.Seed(time.Now().UnixNano())
	//len := len(intArr3)
	//for i:=0;i<len;i++{
	//	intArr3[i] = rand.Intn(100)
	//}
	//fmt.Println(intArr3)
	////反转打印
	//temp :=0
	//for i:=0;i<len/2;i++{
	//	temp = intArr3[i]
	//	intArr3[i] = intArr3[len - i -1]
	//	intArr3[len - i -1] = temp
	//}
	//fmt.Println(intArr3)

	//二维数组

	//定义二位数组
	//var arr [4][6]int
	////fmt.Println(arr)
	//
	//for i:=0;i<len(arr);i++{
	//	for j:=0;j<len(arr[i]);j++{
	//		fmt.Print(" "+strconv.Itoa(arr[i][j]))
	//	}
	//	fmt.Println()
	//}

	//初始化数组
	//var arr = [2][3]int{{1,2,3},{4,5,6}}
	//fmt.Println("arr=",arr)
	//
	//data := []string{"apppcHqlTPT3482","13066867190"}
	//fmt.Println(data)
	//
	////二维数组遍历
	//for i,v := range arr{
	//	for j,v2 := range v{
	//		fmt.Printf("arr[%v][%v]=%v ",i,j,v2)
	//	}
	//	fmt.Println()
	//}
	//
	//	params := genSQLRangeStrByIntArr([]string{"123","13066867190"})
	//	fmt.Println(params)
	//	s :=  fmt.Sprintf("SELECT * FROM t_phone_subscribe_record_0 WHERE (app_id=? AND phone_name=?)" ,params)
	//	fmt.Println(s)
	//}

	//
	//
	//func genSQLRangeStrByIntArr(arr []string) (res string) {
	//	var tempStrArr = make([]string, len(arr))
	//	for k, v := range arr {
	//		tempStrArr[k] = fmt.Sprintf("%s", v)
	//	}
	//	res = "(" + strings.Join(tempStrArr, ",") + ")"
	//	return
	//}
	//


}