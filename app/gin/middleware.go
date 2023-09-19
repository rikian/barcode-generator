package gin

import (
	"log"

	"github.com/gin-gonic/gin"
)

type GinMiddleware struct{}

func (g *GinMiddleware) Middleware(ctx *gin.Context) {
	defer func() {
		r := recover()

		if r != nil {
			log.Print(r)
			ctx.JSON(404, gin.H{
				"msg": "not found",
			})
		}
	}()

	log.Print(ctx.Request.URL)
	ctx.Next()
}
