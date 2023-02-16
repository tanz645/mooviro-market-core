package brandController

import (
	"context"
	"fmt"
	"matar/common/responses"
	"matar/services/brandService"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetBrandById() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		id := c.Param("id")
		defer cancel()
		result, err := brandService.GetBrandById(ctx, id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, responses.FailedResponse{Status: http.StatusNotFound, Error: true, Message: "Can not get Brand", Data: nil})
			return
		}
		data := result
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: data})
	}
}

func GetBrands() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		result, err := brandService.GetBrands(ctx)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Can not get Brands", Data: nil})
			return
		}
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: result})
	}
}
