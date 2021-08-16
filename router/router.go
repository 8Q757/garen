package router

import (
	"github.com/gin-gonic/gin"
	"github.com/viticis/garen/common/net"
	"github.com/viticis/garen/common/net/response"
	"github.com/viticis/garen/router/ping"
)

func Router(r *gin.Engine) {
	noRouter(r)
	ping.Router(r)
}

func noRouter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		response.Fail(net.NotFound, c)
	})
	r.NoMethod(func(c *gin.Context) {
		response.Fail(net.NotFound, c)
	})
}
