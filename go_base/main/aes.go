package main

import (
	"archive/zip"
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 要解压缩的ZIP文件路径
	zipFilePath := "https://static-resource-cos-1252524126.cdn.xiaoeknow.com/xiaoe-cnpm/pro/community-app/1.0.901/offline-package/community_app_1.0.901_release.zip"

	// 本地保存的文件名
	localFileName := "community_app.zip"
	//
	//// 执行文件下载
	err := downloadFile(zipFilePath, localFileName)
	if err != nil {
		fmt.Println("下载文件时发生错误:", err)
		return
	}

	fmt.Printf("文件下载完成，保存在 %s\n", localFileName)

	// 解压缩目标文件夹
	extractFolder := "community_app"

	// 执行解压缩
	unzip(localFileName, extractFolder)
	// 文件路径
	filePath := extractFolder + "/fileList.txt"

	// 读取文件内容
	content, err := readFile(filePath)
	if err != nil {
		fmt.Println("读取文件报错", err)
	}

	// 打印文件内容
	fmt.Println("文件内容:")
	str := string(content)
	strArr := strings.Split(str, "\n")
	fmt.Println("文件列表", strArr)
	strs := ""
	for _, item := range strArr {
		itemContent, err := readFile(extractFolder + "/" + item)
		if item == "manifest.json" {
			manifest := make(map[string]interface{})
			err := json.Unmarshal(itemContent, &manifest)
			if err != nil {
				return
			}
			targetOfflineVersion := manifest["version"].(string)
			fmt.Println("目标更新版本targetOfflineVersion", string(targetOfflineVersion))
		}
		if err != nil {
			continue
		}
		strs += Md5(string(itemContent))
	}
	fmt.Println("文件列表md5组装", string(strs))
	// 16, 24, 或 32 字节的密钥，分别对应 AES-128, AES-192, 或 AES-256
	key := []byte("5cty6xdt7cvtxcdf")
	key2 := []byte("445xtvnb2dkyzxcv")
	plaintext := []byte(Md5(strs))
	fmt.Println("文件列表md5", string(plaintext))

	// 第一次加密
	ciphertext, err := encryptECB(plaintext, key)
	if err != nil {
		fmt.Println("加密时发生错误:", err)
		return
	}

	fmt.Printf("keyA加密后的数据: %s\n", base64.StdEncoding.EncodeToString(ciphertext))

	//第二次加密
	plaintext = []byte(base64.StdEncoding.EncodeToString(ciphertext))
	ciphertextV2, err := encryptECB(plaintext, key2)
	if err != nil {
		fmt.Println("加密时发生错误:", err)
		return
	}
	fmt.Printf("keyB加密后的数据: %s\n", base64.StdEncoding.EncodeToString(ciphertextV2))

	//// 解密
	decryptedText, err := decryptECB(ciphertextV2, key2)
	if err != nil {
		fmt.Println("解密时发生错误:", err)
		return
	}
	txt, _ := base64.StdEncoding.DecodeString(decryptedText)

	decryptedText, err = decryptECB(txt, key)
	fmt.Printf("解密后的数据: %s\n", decryptedText)
}

func downloadFile(url, localFileName string) error {
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，HTTP状态码: %d", resp.StatusCode)
	}

	// 创建本地文件
	file, err := os.Create(localFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将HTTP响应体复制到本地文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func md5Sum(input string) string {
	// 创建MD5哈希对象
	hasher := md5.New()

	// 将字符串转换为字节数组并计算哈希值
	hasher.Write([]byte(input))

	// 获取MD5哈希值的字节数组
	hashBytes := hasher.Sum(nil)

	// 将字节数组转换为十六进制表示
	hashedValue := hex.EncodeToString(hashBytes)

	return hashedValue
}

func readFile(filePath string) ([]byte, error) {
	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// encryptECB 使用AES ECB模式进行加密
func encryptECB(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用PKCS7进行填充
	plaintext = PKCS7Padding(plaintext, block.BlockSize())

	// 加密
	ciphertext := make([]byte, len(plaintext))
	blockSize := block.BlockSize()

	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}

	return ciphertext, nil
}

// decryptECB 使用AES ECB模式进行解密
func decryptECB(ciphertext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 解密
	plaintext := make([]byte, len(ciphertext))
	blockSize := block.BlockSize()

	for i := 0; i < len(ciphertext); i += blockSize {
		block.Decrypt(plaintext[i:i+blockSize], ciphertext[i:i+blockSize])
	}

	// 去除PKCS7填充
	plaintext, err = PKCS7Unpadding(plaintext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// PKCS7Padding 进行PKCS7填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7Unpadding 进行PKCS7解除填充
func PKCS7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("数据为空")
	}

	padding := int(data[length-1])
	if length < padding {
		return nil, errors.New("填充错误")
	}

	return data[:length-padding], nil
}

func unzip(zipFilePath, targetDir string) {

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
