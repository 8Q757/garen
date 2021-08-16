package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/viticis/garen/api/ping"
)

func Router(r *gin.Engine) {
	r.GET("/ping", ping.Ping)
}
