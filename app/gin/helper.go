package gin

import (
	"barcode/models/constants"
	"barcode/models/exceptions"

	g "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewGinHelper(logger *zap.Logger) *GinHelper {
	return &GinHelper{
		logger: logger,
	}
}

type GinHelper struct {
	logger *zap.Logger
}

func (gin *GinHelper) ReadRequestBody(ctx *g.Context, data interface{}) {
	err := ctx.ShouldBindJSON(data)
	if err != nil {
		gin.logger.Error(err.Error())
		panic(exceptions.BadRequest{Msg: constants.InvalidFormatRequest})
	}
}
