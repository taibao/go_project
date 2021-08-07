package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	//创建路由
	//默认使用了2个中间件Logger(),Recovery()
	r := gin.Default()

	//也可以创建不带中间件的路由
	//r := gin.New()

	//2，绑定路由规则，执行的函数
	// gin.Context，封装request和response

	//r.GET( "/",func(c *gin.Context){
	//	c.String(http.StatusOK,"你开什么玩笑？")
	//})
	//r.POST("/XXXPost",getting)
	//r.PUT( "/XXXPut")


	r.GET("/user/:name/*action",func(c *gin.Context){
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK,name+" is "+action)
	})

	//3。监听端口，默认在8080
	r.Run(":8000")
}

func getting(c *gin.Context){

}