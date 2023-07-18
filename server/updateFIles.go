package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func ServerGET5(c *gin.Context) {
	c.HTML(http.StatusOK, "postFiles.html", nil)
}

// ServerPOST5
// 获取数据并且保存到当前目录文件夹下
func ServerPOST5(c *gin.Context) {
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
		})
		log.Fatal(err)
	} else {
		filePath := path.Join("./", file.Filename)
		err1 := c.SaveUploadedFile(file, filePath)
		if err1 != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
