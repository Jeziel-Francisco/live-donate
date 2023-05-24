package infrahttpgin

import (
	"github.com/gin-gonic/gin"
)

func NewServerGinHttp(route RouteGin) *ServerGin {
	return &ServerGin{route: route}
}

type RouteGin interface {
	inicializeRoute(gin *gin.Engine)
}

type ServerGin struct {
	route RouteGin
}

func (s *ServerGin) Run() {
	ginEngine := gin.Default()
	s.route.inicializeRoute(ginEngine)
	ginEngine.Run()
}
