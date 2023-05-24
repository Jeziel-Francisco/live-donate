package main

import (
	corecontroller "github.com/jeziel-francisco/live-donate/core/controller"
	infrahttpgin "github.com/jeziel-francisco/live-donate/infra/http/gin"
)

func main() {
	pingController := corecontroller.NewPingController()
	routeGin := infrahttpgin.NewRouteGin(pingController)
	serverGin := infrahttpgin.NewServerGinHttp(routeGin)
	serverGin.Run()
}
