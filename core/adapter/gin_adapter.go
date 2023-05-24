package coreadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	corecontroller "github.com/jeziel-francisco/live-donate/core/controller"
)

func CreateGin(fn corecontroller.Func) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := c.GetRawData()
		result, _ := fn(body)
		c.JSON(http.StatusOK, result)
	}

}
