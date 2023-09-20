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
			Code:    "AA009995",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "AA009996",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "AA009997",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "AA009998",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "AA009999",
			Printed: true,
			Scanned: false,
			UAt:     time.Now(),
			CAt:     time.Now(),
		},
		{
			NoMes:   "AA",
			Code:    "AA010000",
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
