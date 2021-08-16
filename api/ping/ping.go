package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/viticis/garen/common/net/response"
)

func Ping(c *gin.Context) {
	response.SuccessWithData("pong", c)
}
