package main

import (
	"gin_gorm1/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// gin.Default()
// 默认是使用了Logger和Recovery中间件，其中：
// Logger中间件将日志写入了gin.DefaultWriter，即配置了GIN_MODE=release
// Recovery中间件会recover任何的panic，如果有panic的话就会写入500响应码
var route = gin.Default()

// 可以使用 var route = gin.New() 生成一个空的路由（不含任何中间件）

func main() {

	//渲染模板
	route.LoadHTMLGlob("./templates/*")

	//GET
	route.GET("/login", server.Server1)
	route.GET("/:name/:age", server.Server3)
	route.GET("/login1", server.Server4)
	route.GET("/login2", server.ServerGET4)
	route.GET("/postFiles", server.ServerGET5)
	route.GET("/testR", server.TestRedirect)
	//Test 请求转发，也可以说是内部重定向
	route.GET("/a", func(context *gin.Context) {
		//路由重定向
		context.Request.URL.Path = "/b"
		route.HandleContext(context)
	})
	route.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	//POST
	route.POST("/login", server.Server2)
	route.POST("/test", server.ServerPOST4)
	route.POST("/getFiles", server.ServerPOST5)

	//NoRoute
	route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "找不到界面了",
		})
	})

	//Route Group
	TestGroup := route.Group("/testGroup")
	{
		TestGroup.GET("/page1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "page1",
			})
		})
		TestGroup.GET("/page2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "page2",
			})
		})
		TestGroup.GET("/page3", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "page3",
			})
		})
	}

	//Test Middleware
	route.GET("/middleware", server.MiddlewareTest, server.TestServer)
	//亦可直接全局注册中间件,后续的所有请求都会自动执行该中间件
	//route.Use(server.MiddlewareTest)

	//路由组注册中间件,1
	TestGroup1 := route.Group("/testGroup1", server.AuthMiddleware(true))
	{
		TestGroup1.GET("/page1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "page1",
			})
		})
	}
	//路由组注册中间件,2
	TestGroup2 := route.Group("/testGroup2")
	TestGroup2.Use(server.AuthMiddleware(true))
	{
		TestGroup2.GET("/page1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "page1",
			})
		})
	}

	//RUN
	err := route.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}

}
