package controller

import (
	"gin-template/internal/service"
	"gin-template/pkg/serialize"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	service *service.HelloService
}

func NewHelloController() *HelloController {
	return &HelloController{
		service: &service.HelloService{},
	}
}

func (c *HelloController) Hello(ctx *gin.Context) *serialize.Response {
	msg, code := c.service.GetMessage()
	return serialize.ResponseOK(code, msg)
}
