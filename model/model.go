package model

import "mail_download/model/response_model"

type OrderParam struct {
	FileName     string
	MasterMawbNo string
	QueryInfo    response_model.AirQueryInfo
	PdfData      response_model.ReadPdf
	Login        response_model.LoginResp
}
