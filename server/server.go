package server

import "promotions-app/config"

func Start() {
	config := config.GetConfig()
	router := Router()
	router.Run(":" + config.GetString("server.port"))
}
