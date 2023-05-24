package infrahttpgin

import (
	"github.com/gin-gonic/gin"
	coreadapter "github.com/jeziel-francisco/live-donate/core/adapter"
	corecontroller "github.com/jeziel-francisco/live-donate/core/controller"
)

func NewRouteGin(pingController *corecontroller.PingController) *routeGin {
	return &routeGin{pingController: pingController}
}

type routeGin struct {
	pingController *corecontroller.PingController
}

func (r *routeGin) inicializeRoute(gin *gin.Engine) {
	gin.GET("/ping", coreadapter.CreateGin(r.pingController.Execute))
}
