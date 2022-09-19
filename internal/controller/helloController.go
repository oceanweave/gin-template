package controller

import (
	"gin-template/pkg/serialize"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (h *HelloController) Hello(ctx *gin.Context) *serialize.Response {
	return serialize.NewResponseOk(0, "hello, world")
}
