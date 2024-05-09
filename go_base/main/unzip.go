package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 指定要解压的ZIP文件
	zipFilePath := "community_app.zip"

	// 指定解压目标目录
	targetDir := "community_app"

	unzip1(zipFilePath, targetDir)
}

func unzip1(zipFilePath, targetDir string) {

	// 打开ZIP文件
	zipFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		fmt.Println("无法打开ZIP文件:", err)
		return
	}
	defer zipFile.Close()

	// 遍历ZIP文件中的文件
	for _, file := range zipFile.File {
		// 创建目标文件
		targetFilePath := filepath.Join(targetDir, file.Name)

		// 确保目标文件所在目录存在
		if file.FileInfo().IsDir() {
			// 如果是目录，创建目录
			os.MkdirAll(targetFilePath, os.ModePerm)
			continue
		}

		// 如果是文件，创建文件所在目录
		err := os.MkdirAll(filepath.Dir(targetFilePath), os.ModePerm)
		if err != nil {
			fmt.Println("无法创建目标文件所在目录:", err)
			return
		}

		// 打开ZIP文件中的文件
		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("无法打开ZIP文件中的文件:", err)
			return
		}
		defer fileReader.Close()

		// 创建文件并拷贝数据
		targetFile, err := os.Create(targetFilePath)
		if err != nil {
			fmt.Println("无法创建目标文件:", err)
			return
		}
		defer targetFile.Close()

		// 拷贝数据
		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			fmt.Println("无法拷贝文件数据:", err)
			return
		}
		fmt.Println("解压成功:", targetFilePath)
	}

}
