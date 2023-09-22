package gin

import (
	"barcode/models/exceptions"
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
	defer g.errorResponse(ctx)
	g.logger.Info(ctx.Request.URL.String())
	ctx.Next()
}

func (g *GinMiddleware) errorResponse(ctx *gin.Context) {
	r := recover()
	if r != nil {
		switch err := r.(type) {
		case exceptions.BadRequest:
			log.Print(err.Msg)
			g.logger.Error(err.Msg)
			ctx.JSON(400, gin.H{
				"msg": "bad request",
			})
			break
		case exceptions.NotFound:
			log.Print(err.Msg)
			g.logger.Error(err.Msg)
			ctx.JSON(404, gin.H{
				"msg": "not found",
			})
			break
		case exceptions.ServerError:
			log.Print(err.Msg)
			g.logger.Error(err.Msg)
			ctx.JSON(500, gin.H{
				"msg": "internal server error",
			})
			break
		default:
			v, b := r.(string)
			if b {
				g.logger.Error(v)
			}
			ctx.JSON(500, gin.H{
				"msg": "internal server error",
			})
			log.Print(v)
		}
	}
}
