package main

import (
"fmt"
"github.com/gin-gonic/gin"
"net/http"
)

func main(){
	//创建路由
	//默认使用了2个中间件Logger(),Recovery()


	//也可以创建不带中间件的路由
	//r := gin.New()

	//2，绑定路由规则，执行的函数
	// gin.Context，封装request和response

	//r.GET( "/",func(c *gin.Context){
	//	c.String(http.StatusOK,"你开什么玩笑？")
	//})
	//r.POST("/XXXPost",getting)
	//r.PUT( "/XXXPut")

	//http04()

	//路由组
	routeGroup()
}


//第一种获取参数的方法
func http01(){
	r := gin.Default()
	r.GET("/user/:name/*action",func(c *gin.Context){
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK,name+" is "+action)
	})

	//3。监听端口，默认在8080
	r.Run(":8000")
}

//第二种获取参数的方法
func http02(){
	r := gin.Default()
	r.GET("/welcome",func(c *gin.Context){
		//DefaultQuery
		name := c.DefaultQuery("name","Jack")
		c.String(http.StatusOK,fmt.Sprintf("Hello %s",name))
	})
	//3。监听端口，默认在8080
	r.Run(":8000")
}

//第三种获取表单参数的方法
func http03(){
	r := gin.Default()
	r.POST("/form",func(c *gin.Context){
		//设置表单参数默认值
		type1 :=  c.DefaultPostForm("type","alert")

		//接收参数
		name := c.DefaultPostForm("username","Jack")
		//表单参数设置默认值
		password := c.DefaultPostForm("password","****")
		hobby := c.DefaultPostForm("hobby","run")

		fmt.Println(name)
		if name=="vitas" && password=="123" {
			c.String(http.StatusOK,fmt.Sprintf("Hello %s，你的兴趣是 %s",name,hobby))
		}else{
			c.String(http.StatusOK,fmt.Sprintf("兄得，密码错误！！！"))
			c.String(http.StatusOK,fmt.Sprintf("你的账号是 %s，你的密码是 %s  type是 %s",name,password,type1))
		}

	})

	//3。监听端口，默认在8080
	r.Run(":8000")
}

//上传文件
func http04(){
	r := gin.Default()

	//限制表单上传大小8MB， 默认为32MB
	r.MaxMultipartMemory = 8 << 20

	r.POST("/upload",func(c *gin.Context){
		form,err :=  c.MultipartForm()
		if err != nil{
			c.String(http.StatusOK,fmt.Sprintf("get err "),err.Error())

		}
		//获取所有图片
		files := form.File["avatar"]
		//遍历所有图片

		for _, file := range files{
			//逐个存
			if err := c.SaveUploadedFile(file,file.Filename); err != nil{
				c.String(http.StatusBadRequest,fmt.Sprintf("upload err %s",err.Error()))
				return
			}
		}
		c.String(200,fmt.Sprintf("upload ok %s  files！",len(files)))

		//
		////表单取单个文件
		//file, _ := c.FormFile("avatar")
		//log.Println(file.Filename)
		////传项目根目录，名字为本身
		//c.SaveUploadedFile(file,file.Filename)
		//
		////返回信息
		//c.String(200,fmt.Sprintf("%s upload susuccess！",file.Filename))
		//

	})

	//3。监听端口，默认在8080
	r.Run(":8000")
}


func routeGroup(){
	r := gin.Default()
	//路由组1,处理get请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login",login)
		v1.GET("/submit",submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login",login)
		v2.POST("/submit",submit)
	}
	r.Run(":8000")
}

func login(c *gin.Context){
	name := c.DefaultQuery("name","jack")
	c.String(200,fmt.Sprintf("hello %s\n",name))
}

func submit(c *gin.Context){
	name := c.DefaultPostForm("name","lily")
	c.String(200,fmt.Sprintf("hello %s\n",name))
}