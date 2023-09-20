package gin

import (
	"barcode/models/constants"
	"barcode/service"
	"os"

	g "github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func setupGinEngine() *g.Engine {
	if os.Getenv("ENVIROTMENT") == "PRODUCTION" {
		g.SetMode(g.ReleaseMode)
	}
	gin := g.Default()
	gin.LoadHTMLGlob("./views/*.html")
	return gin
}

func Router(l *zap.Logger, service *service.ServiceBarcode) *g.Engine {
	r := setupGinEngine()
	h := new(GinHelper)
	m := new(GinMiddleware)
	c := new(GinConntroller)
	h.logger, m.logger = l, l
	c.service = service

	// middleware
	r.Use(m.Middleware)

	r.Static("/static/", "./static/")

	// controller
	r.GET(constants.UrlHome, c.Home)
	r.POST(constants.UrlCheckBarcode, c.CheckBarcode)

	return r
}
