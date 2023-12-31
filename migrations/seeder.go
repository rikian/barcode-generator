package migrations

import (
	"barcode/models/entities"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func SeedFakeDataBarcode(dbConn *gorm.DB) {
	tx := dbConn.Begin()
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

	data := []*entities.Barcode{
		{
			NoMes:   "AA",
			Code:    "009995",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "009996",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "009997",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "009998",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "009999",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "010000",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
	}

	sumData := int64(len(data))
	insertData := dbConn.Create(data)

	if insertData.Error != nil {
		panic(insertData.Error.Error())
	}

	if insertData.RowsAffected != sumData {
		panic(fmt.Sprintf("insertData not equal %v", sumData))
	}
}
