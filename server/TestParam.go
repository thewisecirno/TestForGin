package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Server3
// 获取请求参数
func Server3(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
