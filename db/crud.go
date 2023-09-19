package db

import "gorm.io/gorm"

func NewBarcodeImpl(db *gorm.DB) *BarcodeImpl {
	return &BarcodeImpl{
		db: db,
	}
}

type BarcodeImpl struct {
	db *gorm.DB
}

func (c *BarcodeImpl) InsertBarcode() {}
func (c *BarcodeImpl) UpdateBarcode() {}
func (c *BarcodeImpl) DeleteBarcode() {}
func (c *BarcodeImpl) ReadBarcode()   {}
