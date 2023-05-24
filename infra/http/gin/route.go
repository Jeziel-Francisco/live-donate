package infrahttpgin

import (
	"github.com/gin-gonic/gin"
	coreadapter "github.com/jeziel-francisco/live-donate/core/adapter"
	corecontroller "github.com/jeziel-francisco/live-donate/core/controller"
)

func NewRouteGin(pingController, createOrderController corecontroller.Controller) *routeGin {
	return &routeGin{pingController: pingController, createOrderController: createOrderController}
}

type routeGin struct {
	pingController        corecontroller.Controller
	createOrderController corecontroller.Controller
}

func (r *routeGin) inicializeRoute(gin *gin.Engine) {
	gin.GET("/ping", coreadapter.CreateGin(r.pingController.Execute))
	gin.POST("/order", coreadapter.CreateGin(r.createOrderController.Execute))
}
