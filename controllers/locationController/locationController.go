package locationController

import (
	"context"
	"fmt"
	"matar/common/responses"
	"matar/services/locationService"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLocationById() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		id := c.Param("id")
		defer cancel()
		result, err := locationService.GetLocationById(ctx, id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, responses.FailedResponse{Status: http.StatusNotFound, Error: true, Message: "Can not get Location", Data: nil})
			return
		}
		data := result
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: data})
	}
}

func GetLocationsByParentSerial() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		parentSerial := c.Param("parent_serial")
		toInt, err := strconv.Atoi(parentSerial)
		defer cancel()

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Can not get Brands", Data: nil})
			return
		}

		result, err := locationService.GetLocationsByParentSerial(ctx, uint32(toInt))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Can not get Brands", Data: nil})
			return
		}
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: result})
	}
}
