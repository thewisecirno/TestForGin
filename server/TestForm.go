package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Server1(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Server2
// 获取表单数据
func Server2(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, "\n", "password:", password)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Username": username,
		"Password": password,
	})
}
