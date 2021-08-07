package main

import (
	"fmt"
)

//定义列表节点
type HeroNode struct{
	no int
	name string
	nickname string
	pre *HeroNode
	next *HeroNode
}

//双向列表插入节点
func InsertHeroNode_my(head *HeroNode, newHeroNode *HeroNode) {
	//找到链尾,找到适当节点
	//创建辅助节点
	temp := head
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next //让temp不断的指向下一个节点
	}

	temp.next = newHeroNode
	newHeroNode.pre = temp
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
		newHeroNode.next = temp.next //ok
		newHeroNode.pre = temp //ok
		if temp.next != nil{
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}





//显示列表的所有节点信息
func ListHeroNode(head *HeroNode){
	//创建辅助节点
	temp := head

	if temp.next == nil{
		fmt.Println("链表为空")
		return
	}

	for{
		fmt.Printf("[%d,%s,%s]==>",temp.next.no,temp.next.name,temp.next.nickname)
		//判断是否链表后
		temp = temp.next
		if  temp.next == nil{
			break
		}
	}
}

//逆向显示列表的所有节点信息
func ListHeroNode_revert(head *HeroNode){
	//创建辅助节点
	temp := head

	//先判断该列表是否为空
	if temp.next == nil{
		fmt.Println("空的。。。。")
		return
	}

	//temp指向链尾
	for {
		 if temp.next == nil{
			break
		}
		temp = temp.next
	}
	//循环该链表
	for{
		fmt.Printf("[%d,%s,%s]==>",temp.no,temp.name,temp.nickname)
		//判断是否链表头
		temp = temp.pre
		//说明是head
		if  temp.pre == nil{
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
			//编号已存在，不允许插入
			flag = true
			break
		}
		temp =  temp.next
	}
	if flag{
		if temp.next.next != nil{
			temp.next.next.pre = temp
		}
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

	InsertHeroNode_my(head,hero1)
	InsertHeroNode_my(head,hero2)
	InsertHeroNode_my(head,hero3)


	ListHeroNode(head)
	fmt.Println()
	DelNode(head,2)
	DelNode(head,3)
	ListHeroNode_revert(head)
}

