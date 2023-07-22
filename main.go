package main

import (
	"gin_gorm1/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var route = gin.Default()

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

	//RUN
	err := route.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}

}
