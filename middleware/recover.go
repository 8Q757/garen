package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/viticis/garen/common/net"
	"github.com/viticis/garen/common/net/response"
	"log"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			response.Fail(net.InternalServerError, c)
			c.Abort()
		}
	}()
	c.Next()
}
