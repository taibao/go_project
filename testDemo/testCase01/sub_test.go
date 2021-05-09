package main
import (
	//"fmt"
	"testing" // 引入go的testing框架包
)

//编写测试用例 测试GetSub是否正确
func TestGetSub( t *testing.T){
	//调用
	res := getSub(10,8)
	 if res != 2{
	 	//fmt.Println("GetSub(10) 执行错误 期望值 = %v 实际值= %v\n",55 ,res)
	 	t.Fatalf("GetSub(10) 执行错误 期望值 = %v 实际值= %v\n",55 ,res)
	 }

	//如果正确，输出日志
	t.Logf("GetSub(10) 执行正确...")
}
