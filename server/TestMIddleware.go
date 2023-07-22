package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TestServer(c *gin.Context) {
	fmt.Printf("testServer start\n")
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// MiddlewareTest
// hook function
func MiddlewareTest(c *gin.Context) {
	fmt.Println("middlewareTest")
	t := time.Now()
	c.Next()
	diff := time.Since(t)
	fmt.Printf("time:%v\n", diff)
	fmt.Printf("TestServer end")
}
