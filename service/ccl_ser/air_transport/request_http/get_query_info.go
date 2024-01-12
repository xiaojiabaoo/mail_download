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

func GetQueryInfo(login response_model.LoginResp) (response_model.AirQueryInfo, error) {
	var (
		client   = &http.Client{}
		request  = &http.Request{}
		err      error
		post     = &http.Response{}
		response = response_model.AirQueryInfo{}
	)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("url", "/Aim/QueryAimMaster")
	err = writer.Close()
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryInfo接口，关闭writer出现问题")
	}
	request, err = http.NewRequest("POST", "https://us.eftcloud.com/Sss/GetQueryInfo", payload)
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryInfo接口，创建HTTP连接失败")
	}
	request.Header.Set("authority", "us.eftcloud.com")
	request.Header.Set("accept", "*/*")
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("cookie", fmt.Sprintf(`%s; %s`, login.LoginCookie, login.MainCookie))
	request.Header.Set("origin", "https://us.eftcloud.com")
	request.Header.Set("referer", "https://us.eftcloud.com/Efc/MainPage")
	request.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", `"Windows"`)
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	post, err = https.Post(request, client, writer.FormDataContentType())
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryInfo接口，查询数据失败")
	}
	if post == nil || post.Body == nil {
		return response, errors.Wrap(err, "请求GetQueryInfo接口，数据返回失败")
	}
	defer post.Body.Close()
	data, _ := io.ReadAll(post.Body)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryInfo接口，解析数据失败")
	}
	return response, nil
}