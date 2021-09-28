package main

import "fmt"

func main(){

	//演示切片的基本使用
	var intArr [5]int = [...]int{1,22,33,66,99}
	////声明一个切片
	slice := intArr[1:3] //表示引用到intArr的下标1,2的元素。不包括3
	////[22 33]
	//fmt.Println("intArr=",intArr)
	fmt.Println("slice的元素是",slice)
	//fmt.Println("slice的元素个数=",len(slice))
	//fmt.Println("slice的容量=",cap(slice)) //切片的容量是动态变化的，一般是长度的两倍


	//通过make创建切片
	//var slice []float64 = make([]float64,5,10)
	//slice[1] = 10
	//fmt.Println(slice)

	//var strSlice []string = []string{"tom","jack","mary"}
	//fmt.Println(strSlice)
	//fmt.Println()
	//fmt.Println()


	//常规遍历切片
	//var arr [5]int = [...]int{10,20,30,40,50}
	//slice := arr[:]
	//for i := 0;i<len(slice);i++{
	//	fmt.Printf("slice[%v]=%v",i,slice[i])
	//}

	//var slice3 []int = []int{10,20,30}
	////通过append直接给slice3追加具体的元素
	//slice3 = append(slice3,400,500,600)
	//fmt.Println(slice3)
	//
	//slice3 = append(slice3,slice3...)
	//fmt.Println(slice3)
	//
	////copy
	//slice4 := make([]int,10)
	//copy(slice4,slice3)
	//fmt.Println(slice4) //超出下标就会被忽略掉


	//切片求斐波那契数

	//fmt.Println(fbn(10))

}

func fbn(n int)([]uint64){
	fbnslice := make([]uint64,n)
	fbnslice[0] = 1
	fbnslice[1] = 1

	//进行for循环来存放数列
	for i:= 2;i<n;i++{
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}


