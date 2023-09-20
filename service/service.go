package service

import (
	"barcode/db"
	"barcode/models/web"
	"context"

	"go.uber.org/zap"
)

func NewServiceBarcode(logger *zap.Logger, repo *db.BarcodeRepository) *ServiceBarcode {
	return &ServiceBarcode{
		logger: logger,
		repo:   repo,
	}
}

type ServiceBarcode struct {
	logger *zap.Logger
	repo   *db.BarcodeRepository
}

func (s *ServiceBarcode) CheckBarcode(ctx context.Context, req *web.RequestCheckBarcode, res *web.ResponseCheckBarcode) {
	barcodeMatcher := s.repo.CheckBarcode(ctx, req.Data)
	swap := 0

	for i := 0; i < len(req.Data); i++ {
		matcher := false
		for j := swap; j < len(barcodeMatcher); j++ {
			if barcodeMatcher[j].Code == req.Data[i] {
				res.Data[i] = &web.DataResponseCheckBarcode{
					Code:    barcodeMatcher[j].Code,
					Printed: barcodeMatcher[j].Printed,
					Scanned: barcodeMatcher[j].Scanned,
				}

				barcodeMatcher[swap], barcodeMatcher[j] = barcodeMatcher[j], barcodeMatcher[swap]
				swap += 1
				matcher = true
				break
			}
		}

		if matcher {
			continue
		}

		res.Data[i] = &web.DataResponseCheckBarcode{
			Code:    req.Data[i],
			Printed: false,
			Scanned: false,
		}
	}

	res.Code = 200
	res.Msg = "OK"
}
