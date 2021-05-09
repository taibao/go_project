

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//将一个文件内容写入到另一个文件中
func main(){
	srcPath := "o:/1.mp4"
	dstPath := "j:/电影/公牛.mp4"

	_, err:= CopyFile(dstPath,srcPath)
	if err ==nil{
		fmt.Println("拷贝完成")
	}else{
		fmt.Println("拷贝错误 err=%v",err)
	}
}

//接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string ,srcFileName string) (written int64,err error){
	srcFile,err := os.Open(srcFileName)
	if err != nil{
		fmt.Printf("open file err=%v\n",err)
	}
	defer srcFile.Close()
	//通过srcFile，获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile,err := os.OpenFile(dstFileName,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}
	defer dstFile.Close()
	//通过dstFile，获取到Writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer,reader)
}

