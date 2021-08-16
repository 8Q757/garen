package server

import (
	"github.com/gin-gonic/gin"
	"github.com/viticis/garen/middleware"
	"github.com/viticis/garen/router"
)

func InitServer() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middleware.Recover)
	router.Router(r)
	r.Run()
}
