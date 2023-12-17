package request_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mail_download/model/response_model"
	https "mail_download/tools/http"
	"mime/multipart"
	"net/http"
)

func AihHouseExec2(login response_model.LoginResp, charge response_model.AirGetAihChargesRows) (response_model.AirAihHouseExec, error) {
	var (
		client       = &http.Client{}
		request      = &http.Request{}
		err          error
		post         = &http.Response{}
		response     = response_model.AirAihHouseExec{}
		payload      = &bytes.Buffer{}
		writer       = multipart.NewWriter(payload)
		docCodeHouse string
	)
	if charge.DocCodeHouse != nil {
		docCodeHouse = charge.DocCodeHouse.(string)
	}
	_ = writer.WriteField("callsubform", "f_plm4_hbl_list_charge_inbound")
	_ = writer.WriteField("callback", "function () {\n$.clearSelectRowByCk(GetGridId());\n}")
	_ = writer.WriteField("callurl", "/Aih/AihHouseExec")
	_ = writer.WriteField("ask", "ok")
	_ = writer.WriteField("id", fmt.Sprintf(`%d`, charge.PlmHblId))
	_ = writer.WriteField("formid", "41000")
	_ = writer.WriteField("op", "exec")
	_ = writer.WriteField("mode", "Exec")
	_ = writer.WriteField("cmd", "btn0InAPLineInvoiceAir")
	_ = writer.WriteField("precallid", fmt.Sprintf(`%d`, charge.PlmMasterId))
	_ = writer.WriteField("precaller", "AimAimMaster")
	_ = writer.WriteField("caller", "AihAihHouse")
	_ = writer.WriteField("callf", "btn_btn0InAPLineInvoiceAir")
	_ = writer.WriteField("calldoccode", docCodeHouse)
	_ = writer.WriteField("firstgroup", "")
	_ = writer.WriteField("objstr", fmt.Sprintf(`{"ids":[%v],"mid":%v,"group":"Aih","Currency":false}`, charge.KeyId, charge.PlmHblId))
	_ = writer.WriteField("subform", "f_plm4_hbl_list_charge_inbound")
	err = writer.Close()
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，关闭writer出现问题")
	}
	request, err = http.NewRequest("POST", "https://us.eftcloud.com/Aih/AihHouseExec", payload)
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，创建HTTP连接失败")
	}
	cookie := fmt.Sprintf(`%s; %s`, login.LoginCookie, login.MainCookie)
	request.Header.Set("authority", "us.eftcloud.com")
	request.Header.Set("accept", "*/*")
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("cookie", cookie)
	request.Header.Set("origin", "https://us.eftcloud.com")
	request.Header.Set("referer", "https://us.eftcloud.com/Aih/AihHouse?v=v3.1116-09&l=en-US&id=-99&mode=View")
	request.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", `"Windows"`)
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	post, err = https.Post(request, client, writer.FormDataContentType())
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，请求查询数据失败")
	}
	if post == nil || post.Body == nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，请求查询数据返回失败")
	}
	defer post.Body.Close()
	data, _ := io.ReadAll(post.Body)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，请求解析数据失败")
	}
	return response, nil
}

func AihHouseExec(login response_model.LoginResp, charge response_model.AirGetAihChargesRows) (response_model.AirAihHouseExec, error) {
	var (
		client       = &http.Client{}
		request      = &http.Request{}
		err          error
		post         = &http.Response{}
		response     = response_model.AirAihHouseExec{}
		payload      = &bytes.Buffer{}
		writer       = multipart.NewWriter(payload)
		docCodeHouse string
	)
	if charge.DocCodeHouse != nil {
		docCodeHouse = charge.DocCodeHouse.(string)
	}
	_ = writer.WriteField("callsubform", "f_plm4_hbl_list_charge_inbound")
	_ = writer.WriteField("callback", "function () {\n$.clearSelectRowByCk(GetGridId());\n}")
	_ = writer.WriteField("callurl", "/Aih/AihHouseExec")
	_ = writer.WriteField("id", fmt.Sprintf(`%d`, charge.PlmHblId))
	_ = writer.WriteField("mid", fmt.Sprintf(`%d`, charge.PlmMasterId))
	_ = writer.WriteField("op", "exec")
	_ = writer.WriteField("formid", "41000")
	_ = writer.WriteField("mode", "Exec")
	_ = writer.WriteField("cmd", "btn0InAPLineInvoiceAir")
	_ = writer.WriteField("callid", fmt.Sprintf(`%d`, charge.PlmHblId))
	_ = writer.WriteField("caller", "AihAihHouse")
	_ = writer.WriteField("callf", "btn_btn0InAPLineInvoiceAir")
	_ = writer.WriteField("calldoccode", docCodeHouse)
	_ = writer.WriteField("precallid", fmt.Sprintf(`%d`, charge.PlmMasterId))
	_ = writer.WriteField("precaller", "AimAimMaster")
	_ = writer.WriteField("firstgroup", "")
	_ = writer.WriteField("objstr", fmt.Sprintf(`{"ids":[%v],"mid":%v,"group":"Aih","Currency":false}`, charge.KeyId, charge.PlmHblId))
	_ = writer.WriteField("subform", "f_plm4_hbl_list_charge_inbound")
	err = writer.Close()
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，关闭writer出现问题")
	}
	request, err = http.NewRequest("POST", "https://us.eftcloud.com/Aih/AihHouseExec", payload)
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，创建HTTP连接失败")
	}
	cookie := fmt.Sprintf(`%s; %s`, login.LoginCookie, login.MainCookie)
	request.Header.Set("authority", "us.eftcloud.com")
	request.Header.Set("accept", "*/*")
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("cookie", cookie)
	request.Header.Set("origin", "https://us.eftcloud.com")
	request.Header.Set("referer", "https://us.eftcloud.com/Aih/AihHouse?v=v3.1116-09&l=en-US&id=-99&mode=View")
	request.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", `"Windows"`)
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	post, err = https.Post(request, client, writer.FormDataContentType())
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，请求查询数据失败")
	}
	if post == nil || post.Body == nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，请求查询数据返回失败")
	}
	defer post.Body.Close()
	data, _ := io.ReadAll(post.Body)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.Wrap(err, "请求AihHouseExec接口，请求解析数据失败")
	}
	return response, nil
}
