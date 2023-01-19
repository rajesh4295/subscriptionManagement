package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHealthRoute(router *gin.RouterGroup) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Sever is up and running",
		})
	})
}
