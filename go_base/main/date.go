package main

//时间和日期相关函数
func main(){

	//获取当前时间
	//now := time.Now()
	//fmt.Printf("now=%v now type=%T\n",now,now)

	//输出年月日，时分秒
	//fmt.Println("日期=",now.Year(),"-",int(now.Month()),"-",now.Day()," ",now.Hour(),":",now.Minute(),":",now.Second())

	//格式化日期时间
	//fmt.Printf("当前的年月日 %d-%d-%d %d:%d:%d \n",
	//	now.Year(),
	//	int(now.Month()),
	//	now.Day(),
	//	now.Hour(),
	//	now.Minute(),
	//	now.Second())

	//获取格式化后字符串、
	//datestr := fmt.Sprintf("当前的年月日 %d-%d-%d %d:%d:%d \n",
	//	now.Year(),
	//	int(now.Month()),
	//	now.Day(),
	//	now.Hour(),
	//	now.Minute(),
	//	now.Second())
	//
	//fmt.Println(datestr)

	//第二种格式化方式
	//fmt.Println(now.Format("2006/01/02 15:04:05"))
	//fmt.Println(now.Format("2006/01/02"))
	//fmt.Println(now.Format("15:04:05"))

	//time.Sleep(1000 * time.Millisecond)
	//fmt.Println("我等了100毫秒")

	//var i int = 0
	//for{
	//	i++
	//	fmt.Println(i)
	//	time.Sleep(time.Millisecond * 100)
	//	if(i==100){
	//		break
	//	}
	//}

//unix和unixNano的使用
//fmt.Printf("unix时间戳=%v unixnano时间戳=%v",now.Unix(),now.UnixNano())

//start := time.Now().UnixNano()
//test03()
//end := time.Now().UnixNano()
//fmt.Printf("执行test03耗费时间%v纳秒",end-start)





}
//func test03(){
//	str := "bigworld"
//	for i:=0;i<100;i++{
//		fmt.Println(str + strconv.Itoa(i))
//	}
//}