package main

import (
	"fmt"
)

//定义列表节点
type HeroNode struct{
	no int
	name string
	nickname string
	next *HeroNode
}

//插入节点
func InsertHeroNode_my(head *HeroNode,newHeroNode2 *HeroNode) {
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

//插入节点,有序从小到大插入排序
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//找到链尾,找到适当节点
	//创建辅助节点
	temp := head
	flag := true
	//让插入的节点的no，和temp的下一个节点no比较
	for{
		if temp.next == nil{ //说明到列表最后
			break
		}else if temp.next.no < newHeroNode.no{
			//说明newHeroNode就应该插入到temp后面
			break
		}else if  temp.next.no == newHeroNode.no{
			//表号已存在，不允许插入
			flag = false
			break
		}
		temp =  temp.next
	}
	if !flag{
		fmt.Println("已存在no=" , newHeroNode.no)
		return
	}else{
		newHeroNode.next = temp.next
		temp.next =  newHeroNode
	}
}

//显示列表的所有节点信息
func ListHeroNode2(head *HeroNode){
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

func isEmpty(list *HeroNode) bool {
	//判空
	if list.next == nil{
		return true
	}
	return false
}


//删除一个节点
func DelNode(head *HeroNode, id int){
	//创建辅助节点
	temp := head
	flag := false

	//让插入的节点的no，和temp的下一个节点no比较
	for{
		if temp.next == nil{ //说明到列表最后
			break
		}else if  temp.next.no == id{
			//表号已存在，不允许插入
			flag = true
			break
		}
		temp =  temp.next
	}
	if flag{
		temp.next = temp.next.next
	}else{
		fmt.Println("该节点不存在")
	}

}


func main(){
	//创建头节点
	head  := &HeroNode{}

	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
		next:     nil,
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
		next:     nil,
	}


	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
		next:     nil,
	}

	InsertHeroNode2(head,hero1)
	InsertHeroNode2(head,hero3)
	InsertHeroNode2(head,hero2)

	DelNode(head,3)
	DelNode(head,2)
	DelNode(head,1)
	ListHeroNode2(head)
}

