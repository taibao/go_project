package main

import "fmt"

//定义列表节点
type HeroNode2 struct{
	no int
	name string
	nickname string
	next *HeroNode2
}

//插入节点
func InsertHeroNode_my(head *HeroNode2,newHeroNode2 *HeroNode2) {
	//找到链尾,找到适当节点
	//创建辅助节点
	temp := head
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next //让temp不断的指向下一个节点
	}

	temp.next = newHeroNode2

}

//显示列表的所有节点信息
func ListHeroNode2(head *HeroNode2){
	//创建辅助节点
	temp := head
	if isEmpty(temp){
		fmt.Println("链表为空")
		return
	}

	for{
		fmt.Printf("[%d,%s,%s]==>",temp.next.no,temp.next.name,temp.next.nickname)
		//判断是否链表后
		temp = temp.next
		if  isEmpty(temp){
			break
		}
	}
}

func isEmpty(list *HeroNode2) bool {
	//判空
	if list.next == nil{
		return true
	}
	return false
}

func main(){
	//创建头节点
	head  := &HeroNode2{}

	hero1 := &HeroNode2{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
		next:     nil,
	}

	hero2 := &HeroNode2{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
		next:     nil,
	}

	InsertHeroNode_my(head,hero1)
	InsertHeroNode_my(head,hero2)
	ListHeroNode2(head)
}

