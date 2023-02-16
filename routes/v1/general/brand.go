package general

import (
	"matar/controllers/brandController"

	"github.com/gin-gonic/gin"
)

func BrandRoute(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/brands")
	users.GET("/", brandController.GetBrands())
	users.GET("/:id", brandController.GetBrandById())
}
