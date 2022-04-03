package main

import (
	"RuiServer/server/config"
	"RuiServer/server/module/database"
	"RuiServer/server/module/game"
)

func main() {
	config.Init()
	database.Init()
	game.Init()
}
