package gin

import (
	"barcode/models/web"
	"barcode/service"

	"github.com/gin-gonic/gin"
)

type GinConntroller struct {
	helper  *GinHelper
	service *service.ServiceBarcode
}

func (g *GinConntroller) Home(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func (g *GinConntroller) CheckBarcode(ctx *gin.Context) {
	request := &web.RequestCheckBarcode{
		Data: []string{},
	}
	g.helper.ReadRequestBody(ctx, request)
	response := &web.ResponseCheckBarcode{
		Data: make([]*web.DataResponseCheckBarcode, len(request.Data)),
	}
	g.service.CheckBarcode(ctx, request, response)
	ctx.JSON(response.Code, response)
}
