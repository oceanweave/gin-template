package routes

import (
	"gin-template/internal/controller"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	helloController := controller.NewHelloController()
	{
		engine.GET("api/hello", Decorate(helloController.Hello))
	}
}
