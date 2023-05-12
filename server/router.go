package server

import (
	"promotions-app/services/promotions"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	middlewares(router)
	servicesRoutes(router)

	return router

}

func middlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}

func servicesRoutes(router *gin.Engine) {
	promotions.Routes(router)
}
