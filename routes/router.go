package routes

import (
	"gin-template/conf"
	"github.com/gin-gonic/gin"
)

func NewGinRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	var router *gin.Engine
	if conf.ServerConfig.Mode == "dev" {
		router = gin.Default()
	} else {
		router = gin.New()
	}
	if len(middlewares) > 0 {
		router.Use(middlewares...)
	}

	return router
}
