package main

import "fmt"

type CatNode struct{
	no int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode,newCatNode *CatNode){
	//判断是不是添加第一只猫
	if head.next == nil{
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head  //构成一个环形
		fmt.Println(newCatNode,"加入到环形的列表")
		return
	}

	//定义临时变量，找到环形最后节点
	temp := head
	 for{
	 	if temp.next == head{
	 		break
		}
		temp = temp.next
	 }
	 //加入到列表
	 temp.next =newCatNode
	 newCatNode.next = head

}

//输出这个环形连标
func ListCircleLink(head *CatNode){
	temp :=head
	if temp.next ==nil{
		fmt.Println("空的")
		return
	}
	for{
		fmt.Printf("[%d,%s] =>",temp.no,temp.name)
		if temp.next == head{
			break
		}
		temp = temp.next
	}

	fmt.Println()
}


//删除
func DelNode(head *CatNode ,id int) *CatNode{
	temp := head
	helper := head

	//空链表
	if temp.next == nil{
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}

	//如果只有一个节点
	if temp.next == head{
		temp.next = nil
		return head
	}

	//将helper定位到链表最后
	for {
		if helper.next == head{
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		if temp.next == head{
			//如果到这里说明比较到了最后一个
			break
		}
		if temp.no == id{
			if temp == head{
				//说明删除的是头节点
				head = head.next
			}
			helper.next = temp.next
			fmt.Println("猫猫=",id)
			flag = false
			break
		}

		temp = temp.next //用于比较找到该点
		helper = helper.next //用于删除该点

	}

	if flag{
		if temp.no == id{
			helper.next = temp.next
		}else{
			fmt.Println("该节点不存在")
		}
	}
	return head

}


func main(){
	//初始化一个环形链表的头节点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no : 1,
		name : "tom",
	}

	cat2 := &CatNode{
		no : 2,
		name : "lala",
	}


	cat3 := &CatNode{
		no : 3,
		name : "daqiang",
	}


	InsertCatNode(head,cat1)
	InsertCatNode(head,cat2)
	InsertCatNode(head,cat3)
	ListCircleLink(head)
	fmt.Println()
	head = DelNode(head,1)
	fmt.Println()
	ListCircleLink(head)

}