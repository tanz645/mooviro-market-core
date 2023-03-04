package general

import (
	"matar/controllers/locationController"

	"github.com/gin-gonic/gin"
)

func LocationRoute(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/locations")
	users.GET("/", locationController.SearchLocations())
	users.GET("/by-parent/:parent_serial", locationController.GetLocationsByParentSerial())
	users.GET("/:id", locationController.GetLocationById())
}
