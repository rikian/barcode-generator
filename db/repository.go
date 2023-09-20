package db

import (
	"barcode/models/entities"
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewBarcodeRepository(dbConn *gorm.DB, logger *zap.Logger) *BarcodeRepository {
	return &BarcodeRepository{
		logger: logger,
		db:     dbConn,
	}
}

type BarcodeRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (c *BarcodeRepository) CheckBarcode(ctx context.Context, req []string) []*entities.Barcode {
	tx := c.db.Begin()
	if tx.Error != nil {
		panic(tx.Error.Error())
	}

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
			panic(r)
		} else {
			tx.Commit()
		}
	}()

	var dataBarcode []*entities.Barcode
	tx.Where("code IN ?", req).Find(&dataBarcode)
	return dataBarcode
}

func (c *BarcodeRepository) InsertBarcode(ctx context.Context, b ...*entities.Barcode) {
	tx := c.db.Begin()
	if tx.Error != nil {
		panic(tx.Error.Error())
	}

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
			panic(r)
		} else {
			tx.Commit()
		}
	}()

	id := tx.Create(b)

	if id.Error != nil {
		panic(id.Error.Error())
	}

	ra := int64(len(b))

	if id.RowsAffected != ra {
		panic(fmt.Sprintf("insertData not equal %v", ra))
	}
}

func (c *BarcodeRepository) UpdateBarcode() {}
func (c *BarcodeRepository) DeleteBarcode() {}
