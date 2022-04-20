package initialize

import (
	"RuiServer/config"
	"RuiServer/db"
	"RuiServer/router"
)

func Init() {
	config.Init()
	db.Init()
	router.Init()
}
