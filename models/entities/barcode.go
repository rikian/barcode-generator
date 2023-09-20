package entities

import "time"

type Barcode struct {
	ID      uint      `gorm:"not null;unique;primary_key;autoincrement" json:"id"`
	NoMes   string    `gorm:"not null;unique" json:"nomer_mesin"`
	Code    int       `gorm:"not null;unique" json:"code"`
	Printed bool      `gorm:"not null;" json:"printed"`
	Scanned bool      `gorm:"not null;" json:"scanned"`
	UAt     time.Time `gorm:"type:date;size:128;not null" json:"updated_at"`
	CAt     time.Time `gorm:"type:date;size:128;not null" json:"created_at"`
}

type BarcodeHistory struct {
	ID        string    `gorm:"not null;unique;primary_key" json:"id"`
	BarcodeID string    `gorm:"not null;" json:"barcode_id"`
	Barcode   Barcode   `gorm:"foreignkey:BarcodeID;references:ID"`
	Status    string    `gorm:"not null;" json:"status"`
	UAt       time.Time `gorm:"type:date;size:128;not null" json:"updated_at"`
}
