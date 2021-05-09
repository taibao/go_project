package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义结构体，用于保存统计结果
type CharCount struct{
	ChCount    int //记录英文个数
	NumCount   int //记录空格的个数
	SpaceCount int //记录空格数
	OtherCount int //记录其他字符的个数
}

//将一个文件内容写入到另一个文件中
func main(){
	fileName := "1.txt"
	file,err := os.Open(fileName)
	if err != nil{
		fmt.Printf("打开文件错误，err=%v",err)
		return
	}

	defer file.Close()
	//定义个CharCount实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//开始循环的读取fileName的内容
	for{
		str,err := reader.ReadString('\n') //读取一行
		if err == io.EOF{
			break
		}
		//遍历 str，进行统计
		for _,v := range str{
			switch{
				case v >= 'a' && v <= 'z':
					fallthrough
				case v>='A' && v<='Z':
					count.ChCount++
				case v == ' '|| v == '\t':
					count.SpaceCount++
				case v >= '0'|| v <= '9':
					count.NumCount++
				default:
					count.OtherCount++
			}
		}
	}

	fmt.Println(count)

}
