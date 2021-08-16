package core

import (
	"github.com/viticis/garen/core/db"
	"github.com/viticis/garen/core/server"
)

func Start() {
	db.InitDb()
	server.InitServer()
}
