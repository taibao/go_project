package main

import "fmt"

type Boy struct{
	No int
	Next *Boy
}

func AddBoy(num int) *Boy{
	first := &Boy{}

	curBoy := &Boy{}

	if num < 1{
		fmt.Println("num的值错误")
		return first
	}

	//循环构建环形列表
	for i:=1; i<= num; i++{
		boy := &Boy{
			No : i,
		}

		//分析构成循环列表，需要一个辅助指针
		if i == 1{
			first = boy //
			curBoy = boy
			curBoy.Next = first
		} else{
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构成环形列表
		}
	}
	return first
}

//分析思路
//1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int)
//2. 最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {

	//1. 空的链表我们单独的处理
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//留一个，判断 startNO <= 小孩的总数
	//2. 需要定义辅助指针，帮助我们删除小孩
	tail := first
	//3. 让tail执行环形链表的最后一个小孩,这个非常的重要
	//因为tail 在删除小孩时需要使用到.
	for {
		if tail.Next == first { //说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	//4. 让first 移动到 startNo [后面我们删除小孩，就以first为准]
	for i := 1; i <= startNo - 1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	//5. 开始数 countNum, 然后就删除first 指向的小孩
	for {
		//开始数countNum-1次
		for i := 1; i <= countNum -1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		//删除first执行的小孩
		first = first.Next
		tail.Next = first
		//判断如果 tail == first, 圈子中只有一个小孩.
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩小孩编号为%d 出圈 \n", first.No)

}


//显示环形列表
func ShowBoy(first *Boy){

	if first.Next == nil{
		fmt.Println("链表是空的")
		return
	}

	//创建·一个指针帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->",curBoy.No)
		//退出
		if curBoy.Next == first{
			break
		}
		curBoy = curBoy.Next
	}
}

func main(){
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first, 1, 3)
}
