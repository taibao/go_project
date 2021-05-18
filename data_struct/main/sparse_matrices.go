package main

import "fmt"


type ValNode struct{
	row int
	col int
	val int
}

func main(){
	//先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子

	//输出看看原始的数组
	for _,v := range chessMap{
		for _,v2 := range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	//转成稀疏数组 算法
	//（1）.遍历chessMap，如果我们发现一个元素的值不为0，创建一个node结构体
	//(2). 将其放入到对应的切片即可
	var sparseArr []ValNode

	//标准的稀疏数组应该还要记录元素的二维数组规模（行列默认值）
	//创建一个ValNode值接点
	valNode := ValNode{
		row:11,
		col:11,
		val:0,
	}

	sparseArr = append(sparseArr,valNode)
	for i,v := range chessMap{
		for j,v2 := range  v {
			if v2 != 0{
				//创建一个ValNode值结点
				valNode := ValNode{
					row :i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}
	//输出稀疏数组
	for i,valNode := range sparseArr{
		if i==0{
			fmt.Printf("稀疏矩阵共有：%d行 %d列\n",valNode.row,valNode.col)
			continue
		}
		fmt.Printf("第%d个元素:在第%d行 第%d列 值为：%d\n",i,valNode.row,valNode.col,valNode.val)
	}

	//可以将该稀疏数组存盘
	//如何恢复原始数组

	//1.打开d:/chessmap.data =》 恢复原始数组
	//2.这里使用稀疏数组恢复

	//先创建一个原始数组
	var chessMap2 [11][11]int

	//遍历sparseArr【遍历文件每一行】
	for i,valNode := range sparseArr {
		if i != 0{
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}


	//看看chessMap2是不是恢复
	fmt.Println("恢复后的原始数据....")
	for _,v := range chessMap2{
		for _,v2 := range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}



}