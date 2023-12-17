package ccl_ser

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	https "mail_download/tools/http"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Login(account, password string) (response_model.LoginResp, error) {
	var (
		loginClient = &http.Client{}
		mainClient  = &http.Client{}
		loginReq    = &http.Request{}
		mainReq     = &http.Request{}
		err         error
		post        = &http.Response{}
		response    = response_model.LoginResp{}
		param       = request_model.LoginParam{}
		now         = time.Now()
		marshal     []byte
		regex       = regexp.MustCompile(`(\.ASPXAUTH=[^;]+)`)
		get         *http.Response
	)
	param.UserName = account
	param.Password = password
	param.Lang = "zh-CN"
	param.Utc = now.Format("2006-01-02 15:04:05")
	marshal, err = json.Marshal(param)
	if err != nil {
		return response, errors.Wrap(err, "请求登录CLL组装参数失败")
	}
	loginReq, err = http.NewRequest("POST", "https://us.eftcloud.com/Account/EfcLogin", strings.NewReader(string(marshal)))
	if err != nil {
		return response, errors.Wrap(err, "创建HTTP连接失败")
	}
	loginReq.Header.Set("authority", "us.eftcloud.com")
	loginReq.Header.Set("accept", "application/json, text/plain, */*")
	loginReq.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	loginReq.Header.Set("content-type", "application/json;charset=UTF-8")
	loginReq.Header.Set("origin", "https://us.eftcloud.com")
	loginReq.Header.Set("referer", "https://us.eftcloud.com/Account/EfcLogin?ReturnUrl=%2fEfc%2fMainPage")
	loginReq.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	loginReq.Header.Set("sec-ch-ua-mobile", "?0")
	loginReq.Header.Set("sec-ch-ua-platform", `"Windows"`)
	loginReq.Header.Set("sec-fetch-dest", "empty")
	loginReq.Header.Set("sec-fetch-mode", "cors")
	loginReq.Header.Set("sec-fetch-site", "same-origin")
	loginReq.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	post, err = https.Post(loginReq, loginClient, "")
	if err != nil {
		return response, errors.Wrap(err, "请求登录CLL失败")
	}
	if post == nil || post.Body == nil {
		return response, errors.Wrap(err, "请求登录CLL返回数据失败")
	}
	defer post.Body.Close()
	data, _ := io.ReadAll(post.Body)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.Wrap(err, "请求登录CLL解析数据失败")
	}
	response.LoginCookie = regex.FindStringSubmatch(post.Header.Get("Set-Cookie"))[1]
	//---------------------------------------------------------------------
	mainReq, err = http.NewRequest("GET", "https://us.eftcloud.com/Efc/MainPage", nil)
	if err != nil {
		return response, errors.Wrap(err, "创建HTTP连接失败")
	}
	mainReq.Header.Set("authority", "us.eftcloud.com")
	mainReq.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	mainReq.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	mainReq.Header.Set("cache-control", "no-cache")
	mainReq.Header.Set("cookie", response.LoginCookie)
	mainReq.Header.Set("pragma", "no-cache")
	mainReq.Header.Set("referer", "https://us.eftcloud.com/Account/EfcLogin?ReturnUrl=%2fEfc%2fMainPage")
	mainReq.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	mainReq.Header.Set("sec-ch-ua-mobile", "?0")
	mainReq.Header.Set("sec-ch-ua-platform", `"Windows"`)
	mainReq.Header.Set("sec-fetch-dest", "document")
	mainReq.Header.Set("sec-fetch-mode", "navigate")
	mainReq.Header.Set("sec-fetch-site", "same-origin")
	mainReq.Header.Set("sec-fetch-user", "?1")
	mainReq.Header.Set("upgrade-insecure-requests", "1")
	mainReq.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	get, err = https.Get(mainReq, mainClient)
	if err != nil {
		return response, errors.Wrap(err, "请求MainPage失败")
	}
	if get == nil || get.Body == nil {
		return response, errors.Wrap(err, "请求登录CLL返回数据失败")
	}
	trimmedCookie := get.Header.Get("Set-Cookie")
	index := strings.Index(trimmedCookie, ";")
	if index != -1 {
		// 获取第一个分号之前的子字符串
		trimmedCookie = trimmedCookie[:index]
	}
	response.MainCookie = trimmedCookie
	return response, nil
}
