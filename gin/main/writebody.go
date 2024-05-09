package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	router := gin.Default()

	router.Use(ModifyBodyMiddleware())

	router.POST("/user", func(c *gin.Context) {
		// 读取处理修改后的请求体参数
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(500, "Failed to read request body")
			return
		}
		fmt.Println("Modified body:", string(body))
		c.String(200, "Modified body: %s", string(body))
	})

	router.Run(":8080")
}

func ModifyBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取原始请求体参数
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		req := make(map[string]interface{})
		if err != nil {
			c.String(500, "Failed to read request body")
			c.Abort()
			return
		}
		if len(bodyBytes) > 0 {
			err := json.Unmarshal(bodyBytes, &req)
			if err != nil {
				return
			}
		}

		fmt.Println("Original body:", req)
		req["b_user_id"] = "12321312"

		// 修改请求体参数为新的参数
		newBody, _ := json.Marshal(req)

		// 使用修改后的参数创建新的请求体
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))

		// 继续处理下一个中间件或处理器
		c.Next()
	}
}
