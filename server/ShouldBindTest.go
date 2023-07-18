package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Server4(c *gin.Context) {
	var user UserInfo
	//ShouldBind自动匹配对应的字段并且初始化UserInfo对象user
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v \n", user)
	c.JSON(http.StatusOK, user)
}

func ServerGET4(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}

func ServerPOST4(c *gin.Context) {
	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v \n", user)
	c.JSON(http.StatusOK, user)
}
