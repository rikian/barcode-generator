package gin

import (
	"barcode/models/constants"
	"os"

	g "github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

var (
	r *g.Engine
	h *GinHelper
	c *GinConntroller
	m *GinMiddleware
)

func setupGinEngine() *g.Engine {
	if os.Getenv("ENVIROTMENT") == "PRODUCTION" {
		g.SetMode(g.ReleaseMode)
	}
	gin := g.Default()
	gin.LoadHTMLGlob("./views/*.html")
	return gin
}

func Router(l *zap.Logger) *g.Engine {
	r = setupGinEngine()
	h = new(GinHelper)
	m = new(GinMiddleware)
	c = new(GinConntroller)
	h.logger = l

	r.Static("/static/", "./static/")

	// middleware
	r.Use(m.Middleware)

	r.GET(constants.UrlHome, func(ctx *g.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	return r
}
