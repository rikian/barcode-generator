package entities

import "time"

type Barcode struct {
	ID   string    `gorm:"not null;unique;primary_key" json:"id"`
	Code string    `gorm:"not null;unique" json:"code"`
	CAt  time.Time `gorm:"type:date;size:128;not null" json:"created_at"`
}
