package web

import "barcode/models/entities"

type DataBarcode struct {
	ListBarcode [][]string `json:"list_barcode"`
}

type ResponseCheckBarcode struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data []*entities.Barcode `json:"data_barcode"`
}

type ResponseUpdateBarcode struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data []*entities.Barcode `json:"data_barcode"`
}

type ResponseUpdateAvailableBarcode struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data []*entities.Barcode `json:"data_barcode"`
}
