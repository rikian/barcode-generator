package service

import (
	"barcode/db"
	"barcode/models/entities"
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

func (s *ServiceBarcode) CheckBarcode(ctx context.Context, req [][]string, res *web.ResponseCheckBarcode) {
	res.Data = make([]*entities.Barcode, len(req))
	barcodeMatcher := s.repo.CheckBarcode(ctx, req)
	swap := 0
	for i := 0; i < len(req); i++ {
		matcher := false
		for j := swap; j < len(barcodeMatcher); j++ {
			if barcodeMatcher[j].NoMes == req[i][0] && barcodeMatcher[j].Code == req[i][1] {
				res.Data[i] = barcodeMatcher[j]
				barcodeMatcher[swap], barcodeMatcher[j] = barcodeMatcher[j], barcodeMatcher[swap]
				swap += 1
				matcher = true
				break
			}
		}

		if matcher {
			continue
		}

		res.Data[i] = &entities.Barcode{
			NoMes:   req[i][0],
			Code:    req[i][1],
			Printed: false,
			Scanned: false,
		}
	}

	res.Code = 200
	res.Msg = "OK"
}

func (s *ServiceBarcode) UpdateBarcode(ctx context.Context, req [][]string, res *web.ResponseUpdateBarcode) {
	res.Data = nil
	dataBarcodeForInsert := []*entities.Barcode{}
	barcodeMatcher := s.repo.CheckBarcode(ctx, req)
	swap := 0

	for i := 0; i < len(req); i++ {
		matcher := false
		for j := swap; j < len(barcodeMatcher); j++ {
			if barcodeMatcher[j].NoMes == req[i][0] && barcodeMatcher[j].Code == req[i][1] && barcodeMatcher[j].Printed {
				barcodeMatcher[swap], barcodeMatcher[j] = barcodeMatcher[j], barcodeMatcher[swap]
				swap += 1
				matcher = true
				break
			}
		}

		if matcher {
			continue
		}

		dataBarcodeForInsert = append(dataBarcodeForInsert, &entities.Barcode{
			NoMes:   req[i][0],
			Code:    req[i][1],
			Printed: true,
			Scanned: false,
		})
	}

	// insert db for data not matcher
	if len(dataBarcodeForInsert) != 0 {
		s.repo.InsertBarcode(ctx, dataBarcodeForInsert)
	}

	res.Code = 200
	res.Msg = "OK"
}

func (s *ServiceBarcode) UpdateBarcodeAvailableOnly(ctx context.Context, req [][]string, res *web.ResponseUpdateAvailableBarcode) {
	dataBarcodeForInsert := []*entities.Barcode{}
	barcodeMatcher := s.repo.CheckBarcode(ctx, req)
	swap := 0

	for i := 0; i < len(req); i++ {
		matcher := false
		for j := swap; j < len(barcodeMatcher); j++ {
			if barcodeMatcher[j].NoMes == req[i][0] && barcodeMatcher[j].Code == req[i][1] {
				barcodeMatcher[swap], barcodeMatcher[j] = barcodeMatcher[j], barcodeMatcher[swap]
				swap += 1
				matcher = true
				break
			}
		}

		if matcher {
			continue
		}

		// data barcode available for printed
		dataBarcodeForInsert = append(dataBarcodeForInsert, &entities.Barcode{
			NoMes:   req[i][0],
			Code:    req[i][1],
			Printed: true,
			Scanned: false,
		})
	}

	// insert db for data not matcher
	if len(dataBarcodeForInsert) != 0 {
		s.repo.InsertBarcode(ctx, dataBarcodeForInsert)
	}

	res.Data = dataBarcodeForInsert
	res.Code = 200
	res.Msg = "OK"
}
