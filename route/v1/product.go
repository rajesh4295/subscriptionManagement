package route

import (
	"net/http"
	"strconv"
	"subscriptionManagement/dao"
	"subscriptionManagement/model"
	"subscriptionManagement/service"

	"github.com/gin-gonic/gin"
)

func InitProductRoute(router *gin.RouterGroup, appCtx *model.AppCtx) {
	productRoute := router.Group("/product")

	handler := &dao.Handler{
		ProductDAO: &service.Product{
			DB: appCtx.DB,
		},
	}

	productRoute.GET("/", func(ctx *gin.Context) {
		var products []model.Product
		if result := handler.ProductDAO.Get(&products); result.Error != nil {
			ErrorHandler(ctx, http.StatusInternalServerError, result.Error.Error())
		} else {
			ResponseHandler(ctx, http.StatusOK, products, "")
		}
	})

	productRoute.GET("/:id", func(ctx *gin.Context) {
		var product model.Product
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		result := handler.ProductDAO.GetById(uint(id), &product)
		if result.Error != nil {
			ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if result.RowsAffected > 0 {
			ResponseHandler(ctx, http.StatusOK, product, "")
		} else {
			ErrorHandler(ctx, http.StatusNotFound, "Product id not found")
		}
	})

	productRoute.POST("/", func(ctx *gin.Context) {
		var product model.Product

		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
		} else {
			if result := handler.ProductDAO.Create(&product); result.Error != nil {
				ErrorHandler(ctx, http.StatusInternalServerError, result.Error.Error())
			} else {
				ResponseHandler(ctx, http.StatusOK, product, "Product created successfully")
			}
		}
	})

	productRoute.PUT("/", func(ctx *gin.Context) {
		var product model.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
		} else {
			var searchProduct model.Product
			result := handler.ProductDAO.GetById(product.ID, &searchProduct)
			if result.Error != nil {
				ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
			} else if result.RowsAffected > 0 {
				if result = handler.ProductDAO.Update(&product); result.Error != nil {
					ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
				} else {
					ResponseHandler(ctx, http.StatusOK, product, "Product updated successfully")
				}
			} else {
				ErrorHandler(ctx, http.StatusNotFound, "Product id not found")
			}
		}
	})

	productRoute.DELETE("/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
		} else {
			result := handler.ProductDAO.Delete(uint(id))
			if result.Error != nil {
				ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
			} else if result.RowsAffected > 0 {
				ResponseHandler(ctx, http.StatusOK, nil, "Product deleted successfully")
			} else {
				ErrorHandler(ctx, http.StatusNotFound, "Product id not found")
			}
		}
	})

}
func ResponseHandler(ctx *gin.Context, statusCode int, data interface{}, message string) {
	ctx.JSON(statusCode, gin.H{
		"data":    data,
		"message": message,
	})
}

func ErrorHandler(ctx *gin.Context, statusCode int, err string) {
	ctx.JSON(statusCode, gin.H{
		"error": err,
	})
}
