package v1

import (
	"gingonic-test/product"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	//router.Use(controllers.Cors())
	// v1 of the API
	v1 := router.Group("/api/v1")
	{
		v1.GET("/product", product.GetAllProduct)
		v1.GET("/product/:id", product.GetProduct)
		v1.POST("/product", product.SaveProduct)
		v1.DELETE("/product/:id", product.DeleteProduct)
		v1.PUT("/product/:id", product.UpdateProduct)
	}
}
