package routes

import (
	"gin-template/pkg/serialize"
	"github.com/gin-gonic/gin"
)

type Handler func(ctx *gin.Context) *serialize.Response

func Decorate(h Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := h(ctx)
		if r != nil {
			ctx.JSON(r.HttpStatus, &r.R)
		}

		serialize.PutResponse(r)
	}
}
