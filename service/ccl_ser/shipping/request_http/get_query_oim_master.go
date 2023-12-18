package request_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"mail_download/model/response_model"
	https "mail_download/tools/http"
	"mime/multipart"
	"net/http"
)

func GetQueryOimMaster(info response_model.AirQueryInfo, masterMawbNo string, login response_model.LoginResp) (response_model.ShiQueryOimMaster, error) {
	var (
		client   = &http.Client{}
		request  = &http.Request{}
		err      error
		post     = &http.Response{}
		response = response_model.ShiQueryOimMaster{}
		payload  = &bytes.Buffer{}
		writer   = multipart.NewWriter(payload)
		vid      string
	)
	if info.Vid != nil {
		vid = info.Vid.(string)
	}
	_ = writer.WriteField("caller", "refreshdg_byfilter_nofooter")
	_ = writer.WriteField("vid", vid)
	_ = writer.WriteField("filterStr", fmt.Sprintf(`[{"type":"string","field":"master_bl_no","value":"%s","comparison":"eq"}]`, masterMawbNo))
	_ = writer.WriteField("action", "Query")
	_ = writer.WriteField("type", "byfilter_nofooter")
	_ = writer.WriteField("efcrows", "25")
	_ = writer.WriteField("page", "1")
	_ = writer.WriteField("rows", "20")
	_ = writer.WriteField("sort", "__")
	_ = writer.WriteField("order", "__desc")
	err = writer.Close()
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryOimMaster接口，关闭writer出现问题")
	}
	request, err = http.NewRequest("POST", "https://us.eftcloud.com/Oim/GetQueryOimMaster", payload)
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryOimMaster接口，创建HTTP连接失败")
	}
	request.Header.Set("authority", "us.eftcloud.com")
	request.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("cookie", fmt.Sprintf(`%s; %s`, login.LoginCookie, login.MainCookie))
	request.Header.Set("origin", "https://us.eftcloud.com")
	request.Header.Set("referer", fmt.Sprintf("https://us.eftcloud.com/Oim/QueryOimMaster?vid=%v&qv=%d&l=en-US&op=get&mode=Query&callid=0", vid, info.Ver))
	request.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", `"Windows"`)
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	post, err = https.Post(request, client, writer.FormDataContentType())
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryOimMaster接口，查询数据失败")
	}
	if post == nil || post.Body == nil {
		return response, errors.Wrap(err, "请求GetQueryOimMaster接口，数据返回失败")
	}
	defer post.Body.Close()
	data, _ := ioutil.ReadAll(post.Body)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.Wrap(err, "请求GetQueryOimMaster接口，解析数据失败")
	}
	return response, err
}
