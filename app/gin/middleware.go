package gin

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewGinMiddleware(logger *zap.Logger) *GinMiddleware {
	return &GinMiddleware{
		logger: logger,
	}
}

type GinMiddleware struct {
	logger *zap.Logger
}

func (g *GinMiddleware) Middleware(ctx *gin.Context) {
	defer func() {
		r := recover()

		if r != nil {
			v, _ := r.(string)
			g.logger.Error(v)
			log.Print(r)
			ctx.JSON(404, gin.H{
				"msg": "not found",
			})
		}
	}()

	g.logger.Info(ctx.Request.URL.String())
	ctx.Next()
}
