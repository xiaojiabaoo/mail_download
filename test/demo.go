package main

import (
	"fmt"
	"mail_download/model/response_model"
	"mail_download/service/ccl_ser"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	var (
		err     error
		fileMap = map[string]response_model.ReadPdf{}
	)
	fileMap, err = ccl_ser.ReadPdfGroup("D:\\ActiveFile\\CCL\\文件")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for k, v := range fileMap {
		fmt.Println("=====================")
		fmt.Println(k)
		fmt.Println(v.Amount)
		fmt.Println(v.MasterMawbNo)
		fmt.Println(v.InvoiceNo)
	}
}
