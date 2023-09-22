package gin

import (
	"barcode/models/web"
	"barcode/service"

	"github.com/gin-gonic/gin"
)

func NewGinConntroller(helper *GinHelper, service *service.ServiceBarcode) *GinConntroller {
	return &GinConntroller{
		helper:  helper,
		service: service,
	}
}

type GinConntroller struct {
	helper  *GinHelper
	service *service.ServiceBarcode
}

func (g *GinConntroller) Home(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func (g *GinConntroller) CheckBarcode(ctx *gin.Context) {
	request := &web.DataBarcode{}
	g.helper.ReadRequestBody(ctx, request)
	response := &web.ResponseCheckBarcode{}
	g.service.CheckBarcode(ctx, request.ListBarcode, response)
	ctx.JSON(response.Code, response)
}

func (g *GinConntroller) UpdateBarcode(ctx *gin.Context) {
	request := &web.DataBarcode{}
	g.helper.ReadRequestBody(ctx, request)
	response := &web.ResponseUpdateBarcode{}
	g.service.UpdateBarcode(ctx, request.ListBarcode, response)
	ctx.JSON(response.Code, response)
}

func (g *GinConntroller) UpdateAvailableBarcode(ctx *gin.Context) {
	request := &web.DataBarcode{}
	g.helper.ReadRequestBody(ctx, request)
	response := &web.ResponseUpdateAvailableBarcode{}
	g.service.UpdateBarcodeAvailableOnly(ctx, request.ListBarcode, response)
	ctx.JSON(response.Code, response)
}
