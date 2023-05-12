package promotions

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	routesGroup := router.Group("promotions")
	controller := new(Controller)

	routesGroup.GET("/:id", controller.Find)
	routesGroup.POST("/", controller.Upload)
}
