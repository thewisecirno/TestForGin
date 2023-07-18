package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func S1(c *gin.Context) {
	type User struct {
		Name string
		Age  int
	}
	user := User{
		Name: "xiaoming",
		Age:  18,
	}

	data := map[string]interface{}{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}

	c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, data)
}
