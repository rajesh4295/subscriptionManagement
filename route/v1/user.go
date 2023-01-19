package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(router *gin.RouterGroup) {
	userRoute := router.Group("/user")

	userRoute.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "this is user route",
		})
	})
}
