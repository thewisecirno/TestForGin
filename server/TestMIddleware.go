package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TestServer(c *gin.Context) {
	fmt.Printf("testServer start\n")

	//取上下文c中的name值
	name, ok := c.Get("name")
	if !ok {
		name = "Not Found"
	}
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

// MiddlewareTest
// hook function
func MiddlewareTest(c *gin.Context) {
	fmt.Println("middlewareTest")
	t := time.Now()
	//中间件中使用goroutine的时候不能直接使用c，需要使用c的拷贝副本
	// go function(c.Copy())
	//给c中赋值
	c.Set("name", "Cirno")
	c.Next() //继续后续的函数调用
	// c.Abort() 阻止后续的函数调用
	diff := time.Since(t)
	fmt.Printf("time:%v\n", diff)
	fmt.Printf("TestServer end")
}

// AuthMiddleware
// 实际开发中，中间件一般会采取闭包的形式
func AuthMiddleware(ok bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ok {
			/*
				存放具体的逻辑
				if 是登陆用户
					c.Next()
				else
					c.Abort()
			*/
			c.Next()
		} else {
			c.Next()
		}
	}
}
