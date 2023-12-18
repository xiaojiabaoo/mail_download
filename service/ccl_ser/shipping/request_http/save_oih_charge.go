package request_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/tools"
	https "mail_download/tools/http"
	"mime/multipart"
	"net/http"
)

func SaveShiChargeJsonParam(charges response_model.ShiGetOihChargesRows) request_model.SaveOihChargeParam {
	var (
		param = request_model.SaveOihChargeParam{}
	)
	param.TableKeyName = tools.AnyToString(charges.TableKeyName)
	param.IsFooter = charges.IsFooter
	param.Selector = charges.Selector
	param.KeyId = charges.KeyId
	param.MasterId = charges.MasterId
	param.Caller = charges.Caller
	param.PlmChargeCostId = charges.PlmChargeCostId
	param.EfcCompanyId = charges.EfcCompanyId
	param.EfcBranchId = charges.EfcBranchId
	param.PlmChargeCostSeqNo = charges.PlmChargeCostSeqNo
	param.BranchCode = tools.AnyToString(charges.BranchCode)
	param.PliChargeCostItemId = charges.PliChargeCostItemId
	param.PliChargeCostItemNo = fmt.Sprintf(`%d`, charges.PliChargeCostItemNo)
	param.OpType = tools.AnyToString(charges.OpType)
	param.ChargeCostType = tools.AnyToString(charges.ChargeCostType)
	param.PlmCodeIndicator = tools.AnyToString(charges.PlmCodeIndicator)
	param.PlmCatCode = tools.AnyToString(charges.PlmCatCode)
	param.PlmMasterId = charges.PlmMasterId
	param.PlmHblId = charges.PlmHblId
	param.PlmHblDisplaySequence = tools.AnyToString(charges.PlmHblDisplaySequence)
	param.DocCodeMaster = tools.AnyToString(charges.DocCodeMaster)
	param.DocCodeHouse = tools.AnyToString(charges.DocCodeHouse)
	param.PrepaidCharge = charges.PrepaidCharge
	param.CollectCharge = charges.CollectCharge
	param.ArCharge = fmt.Sprintf(`%v`, charges.ArCharge)
	param.PrepaidCost = charges.PrepaidCost
	param.CollectCost = charges.CollectCost
	param.ApCost = fmt.Sprintf(`%v`, charges.ApCost)
	param.CurrencySw = charges.CurrencySw
	param.PrepaidChargeCurrency = charges.PrepaidChargeCurrency
	param.CollectChargeCurrency = charges.CollectChargeCurrency
	param.ArChargeCurrency = charges.ArChargeCurrency
	param.PrepaidCostCurrency = charges.PrepaidCostCurrency
	param.CollectCostCurrency = charges.CollectCostCurrency
	param.ApCostCurrency = charges.ApCostCurrency
	param.CurrencyCode = charges.CurrencyCode
	param.CurrencyRate = charges.CurrencyRate
	param.CurrencyXDCode = charges.CurrencyXDCode
	param.ArChargeOtherCurrency = charges.ArChargeOtherCurrency
	param.ApCostOtherCurrency = charges.ApCostOtherCurrency
	param.Reference = charges.Reference
	param.ExtraDesc = tools.AnyToString(charges.ExtraDesc)
	param.VendorId = charges.VendorId
	param.CustomerId = charges.CustomerId
	param.ArBilledSw = charges.ArBilledSw
	param.ApBilledSw = charges.ApBilledSw
	param.AgentCalcSw = charges.AgentCalcSw
	param.ExtraInvSw = charges.ExtraInvSw
	param.ExtraCostSw = charges.ExtraCostSw
	param.DueAgentSw = charges.DueAgentSw
	param.ArmInvoiceId = charges.ArmInvoiceId
	param.ArmInvoiceDetailId = charges.ArmInvoiceDetailId
	param.ApiInvoiceId = charges.ApiInvoiceId
	param.ApiInvoiceDetailId = charges.ApiInvoiceDetailId
	param.RateFactorTotal = charges.RateFactorTotal
	param.RateUnitCode = charges.RateUnitCode
	param.RateAmt = charges.RateAmt
	param.CalculateAmt = charges.CalculateAmt
	param.FrtDistributeCost = charges.FrtDistributeCost
	param.VendorName = "CCL GROUP IN - USLAX"
	param.CustomerName = tools.AnyToString(charges.CustomerName)
	param.ItemCharge = charges.ItemCharge
	param.ItemCost = charges.ItemCost
	param.VatTaxSw = charges.VatTaxSw
	param.ApTaxPercent = charges.ApTaxPercent
	param.ApTaxAmt = charges.ApTaxAmt
	param.ApWithholdTaxPercent = charges.ApWithholdTaxPercent
	param.ApWithholdTaxAmt = charges.ApWithholdTaxAmt
	param.ArWithholdTaxPercent = charges.ArWithholdTaxPercent
	param.ArWithholdTaxAmt = charges.ArWithholdTaxAmt
	param.ArTaxPercent = charges.ArTaxPercent
	param.ArTaxAmt = charges.ArTaxAmt
	param.ApTaxCurrency = charges.ApTaxCurrency
	param.ArTaxCurrency = charges.ArTaxCurrency
	param.PeriodId = charges.PeriodId
	param.AgentArId = charges.AgentArId
	param.AgentApId = charges.AgentApId
	param.AgentName = charges.AgentName
	param.OldPlmChargeCostId = charges.OldPlmChargeCostId
	param.OldPlmBranchCode = charges.OldPlmBranchCode
	param.PlmDisplayFloatSeq = charges.PlmDisplayFloatSeq
	param.CreatedDate = tools.AnyToString(charges.CreatedDate)
	param.CreatedByCode = tools.AnyToString(charges.CreatedByCode)
	param.ModifiedDate = tools.AnyToString(charges.ModifiedDate)
	param.ModifiedByCode = tools.AnyToString(charges.ModifiedByCode)
	param.Rowversion = charges.Rowversion
	param.AdvanceItemSw = charges.AdvanceItemSw
	param.ChargeCostDesc = tools.AnyToString(charges.ChargeCostDesc)
	param.SelSw = charges.SelSw
	param.X12ChargeCode = charges.X12ChargeCode
	param.ObjectState = charges.ObjectState
	param.Selector1 = ""
	param.Color = ""
	param.DbAction = 0
	return param
}

func SaveOihCharge(login response_model.LoginResp, charges response_model.ShiGetOihChargesRows) (response_model.ShiSaveOihCharge, error) {
	var (
		client   = &http.Client{}
		request  = &http.Request{}
		err      error
		post     = &http.Response{}
		response = response_model.ShiSaveOihCharge{}
		param    = request_model.SaveOihChargeParam{}
		payload  = &bytes.Buffer{}
		writer   = multipart.NewWriter(payload)
		marshal  []byte
	)
	param = SaveShiChargeJsonParam(charges)
	marshal, err = json.Marshal(param)
	if err != nil {
		return response, errors.Wrap(err, "序列化SaveOihCharge参数失败")
	}
	_ = writer.WriteField("jsonStr", string(marshal))
	_ = writer.WriteField("needrow", "true")
	_ = writer.WriteField("id", fmt.Sprintf(`%v`, param.KeyId))
	_ = writer.WriteField("mid", "-1")
	_ = writer.WriteField("op", "save")
	_ = writer.WriteField("mode", "Edit")
	_ = writer.WriteField("ask", "")
	_ = writer.WriteField("edittype", "inline")
	err = writer.Close()
	if err != nil {
		return response, errors.Wrap(err, "请求SaveOihCharge接口，关闭writer出现问题")
	}
	request, err = http.NewRequest("POST", "https://us.eftcloud.com/Oih/SaveOihCharge", payload)
	if err != nil {
		return response, errors.Wrap(err, "请求SaveOihCharge接口，创建HTTP连接失败")
	}
	cookie := fmt.Sprintf(`%s; %s`, login.LoginCookie, login.MainCookie)
	request.Header.Set("authority", "us.eftcloud.com")
	request.Header.Set("accept", "*/*")
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("cookie", cookie)
	request.Header.Set("origin", "https://us.eftcloud.com")
	request.Header.Set("referer", "https://us.eftcloud.com/Oih/OihHouse?v=v3.1116-09&l=en-US&id=-99&mode=View")
	request.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", `"Windows"`)
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	post, err = https.Post(request, client, writer.FormDataContentType())
	if err != nil {
		return response, errors.Wrap(err, "请求SaveOihCharge接口，请求查询数据失败")
	}
	if post == nil || post.Body == nil {
		return response, errors.Wrap(err, "请求SaveOihCharge接口，请求查询数据返回失败")
	}
	defer post.Body.Close()
	data, _ := io.ReadAll(post.Body)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, errors.Wrap(err, "请求SaveOihCharge接口，请求解析数据失败")
	}
	return response, nil
}
