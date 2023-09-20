package web

type RequestCheckBarcode struct {
	Data []string `json:"data_barcode"`
}

type DataResponseCheckBarcode struct {
	Code    string `gorm:"not null;unique" json:"code"`
	Printed bool   `gorm:"not null;" json:"printed"`
	Scanned bool   `gorm:"not null;" json:"scanned"`
}

type ResponseCheckBarcode struct {
	Code int                         `json:"code"`
	Msg  string                      `json:"msg"`
	Data []*DataResponseCheckBarcode `json:"data_barcode"`
}
