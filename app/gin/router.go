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

func Router(l *zap.Logger, s *service.ServiceBarcode) *g.Engine {
	r := setupGinEngine()
	h := NewGinHelper(l)
	m := NewGinMiddleware(l)
	c := NewGinConntroller(h, s)

	// middleware
	r.Use(m.Middleware)

	// serve static
	r.Static("/static/", "./static/")

	// controller
	r.GET(constants.UrlHome, c.Home)
	r.POST(constants.UrlCheckBarcode, c.CheckBarcode)
	r.POST(constants.UrlUpdateBarcode, c.UpdateBarcode)
	r.POST(constants.UrlAvailableBarcode, c.UpdateAvailableBarcode)

	return r
}
