package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ModifyResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里执行中间件的其他逻辑...

		// 构建一个 User 结构体实例
		user := User{
			Username: "example_user",
			Email:    "example_user@example.com",
		}

		// 将 User 结构体序列化为 JSON 格式
		userJSON, err := json.Marshal(user)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// 设置响应体内容类型和写入 JSON 数据
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(userJSON)

		// 终止请求处理链，因为我们已经写入了响应体
		c.Abort()
	}
}

func main() {
	r := gin.Default()

	// 使用中间件
	r.Use(ModifyResponseMiddleware())

	r.GET("/user", func(c *gin.Context) {
		// 由于中间件已经写入了响应体并终止了处理链，这个路由处理函数将不会被执行
	})

	// 启动服务器
	r.Run(":8080")
}