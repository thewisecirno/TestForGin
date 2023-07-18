package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TestRedirect
// 内部重定向（重写）发生在server内部，
// client不知情，即浏览器上显示的是访问a.html但实际上收到的却是b.php；
// 外部重定向则是server端通知client端需要更改URL，并舍弃之前的request。
// client按要求发送第二个request，server发送对应的文件。
// 此处可以视为是外部重定向
func TestRedirect(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://www.bing.com")
}
