package model

import "mail_download/model/response_model"

type OrderParam struct {
	FileName     string
	MasterMawbNo string
	QueryInfo    response_model.AirQueryInfo
	PdfData      response_model.ReadPdf
	Login        response_model.LoginResp
}

type XlsxData struct {
	Time     string `json:"time"`     // 操作时间
	Level    string `json:"level"`    // 记录等级
	Describe string `json:"describe"` // 记录描述
	Response string `json:"response"` // 操作第三方响应数据
}
