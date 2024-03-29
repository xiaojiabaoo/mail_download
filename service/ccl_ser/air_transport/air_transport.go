package air_transport

import (
	"fmt"
	"mail_download/model"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/service/ccl_ser/air_transport/request_http"
	"mail_download/tools"
)

func AirTransport(order model.OrderParam, param request_model.CCLParam) (int64, error) {
	var (
		err         error
		master      response_model.AirQueryAimMaster
		hbls        response_model.AirGetAimHbls
		charges     response_model.AirGetAihCharges
		chargesRows response_model.AirGetAihChargesRows
		dutyCount   int64
		charge      response_model.AirSaveAihCharge
		exec        response_model.AirAihHouseExec
		apis        response_model.AirGetAimReviewInvoiceApis
		types       int64 = 1 // 1.表示流程因数据原因操作失败（包括因系统/第三方错误导致的失败）；2.表示操作成功；
	)
	//-------------------请求QueryAimMaster数据-------------------
	//获取表格数据：网站设计原因 这里要请求两次
	tools.Logger(param.Serial, "开始查询AI MAWB列表，查询主单号（MAWB No）："+order.MasterMawbNo, "", "")
	for i := 0; i < 2; i++ {
		master, err = request_http.GetQueryAimMaster(order.QueryInfo, order.MasterMawbNo, order.Login)
		if err != nil {
			return types, err
		}
		if len(master.Rows) > 0 {
			break // 如果查询到了数据：直接跳过，不再继续查询
		}
	}
	switch {
	case len(master.Rows) == 0:
		tools.Logger(param.Serial, "AI MAWB中未检测到数据，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	case len(master.Rows) > 1:
		tools.Logger(param.Serial, "检测到AI MAWB中有多条数据，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	case master.Rows[0].StatusDesc == "Close":
		tools.Logger(param.Serial, "检测到AI MAWB中的数据状态为：Close，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	case master.Rows[0].StatusDesc == "Open":
		tools.Logger(param.Serial, "检测到AI MAWB中的数据状态为：Open，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	case master.Rows[0].MasterBlNo != order.MasterMawbNo:
		tools.Logger(param.Serial, "检测到AI MAWB中主单号和表格中的”MAWB No“列不一致，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	}
	tools.Logger(param.Serial, "查询AI MAWB列表完成", fmt.Sprintf(`查询到的数据信息：%+v`, master.Rows[0]), "")
	//-------------------请求GetAimReviewInvoiceApis-------------------
	tools.Logger(param.Serial, "模拟点击当前主单号，跳转A/R-A/P检测是否已生成AP数据", "", "")
	apis, err = request_http.GetAimReviewInvoiceApis(master.Rows[0], order.Login)
	if err != nil {
		return types, err
	}
	for _, v := range apis.Rows {
		if v.PayeeName == "CCL GROUP INC" {
			tools.Logger(param.Serial, "已生成AP数据，请人工核实", "", tools.LOG_LEVEL_ERROR)
			return types, nil
		}
	}
	tools.Logger(param.Serial, "检测完成，未生成AP数据", "", "")
	//-------------------请求GetAimHbls数据-------------------
	tools.Logger(param.Serial, "前往”List HAWBs“页面", "", "")
	hbls, err = request_http.GetAimHbls(master.Rows[0].AimMasterId, order.Login)
	if err != nil {
		return types, err
	}
	if len(hbls.Rows) == 0 {
		tools.Logger(param.Serial, "List HAWBs中未检测到数据，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	}
	if len(hbls.Rows) > 1 {
		tools.Logger(param.Serial, "检测到List HAWBs中数据为一主多分，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	}
	tools.Logger(param.Serial, "查询List HAWBs列表完成", fmt.Sprintf(`查询到的数据信息：%+v`, hbls.Rows[0]), "")
	//-------------------请求GetAihCharges数据-------------------
	tools.Logger(param.Serial, "模拟点击”List HAWBs“中为当前主单号的数据，跳转至”Charge“页面", "", "")
	charges, err = request_http.GetAihCharges(hbls.Rows[0].AihHblId, order.Login)
	if err != nil {
		return types, err
	}
	if len(charges.Rows) == 0 {
		tools.Logger(param.Serial, "Charge中未检测到数据，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	}
	for _, v := range charges.Rows {
		if v.ChargeCostDesc == "Duty" {
			chargesRows = v
			dutyCount++
		}
	}
	switch {
	case dutyCount > 1:
		tools.Logger(param.Serial, "检测到Charge中Charge Item为Duty的数量大于一条，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	case dutyCount == 0:
		tools.Logger(param.Serial, "检测到Charge中Charge Item为Duty的数量为0，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	case chargesRows.ArCharge != chargesRows.ApCost:
		tools.Logger(param.Serial, "检测到Charge中收付不一致，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	case chargesRows.ApCost != tools.StringToFloat64(order.PdfData.Amount):
		tools.Logger(param.Serial, "检测到Charge中应付成本与实际成本不一致，请人工核实", "", tools.LOG_LEVEL_ERROR)
		return types, nil
	}
	tools.Logger(param.Serial, "查询Charge列表完成", fmt.Sprintf(`查询到的数据信息：%+v`, charges.Rows), "")
	//-------------------保存Duty数据-------------------
	tools.Logger(param.Serial, "开始保存Duty数据，在”Vendor“中填写：CCL GROUP IN - USLAX，然后再在”Ref-Pro No”中填写："+order.PdfData.InvoiceNo, "", "")
	chargesRows.Reference = order.PdfData.InvoiceNo
	charge, err = request_http.SaveAihCharge(order.Login, chargesRows)
	if err != nil {
		return types, err
	}
	tools.Logger(param.Serial, "保存Duty数据完成", fmt.Sprintf(`系统响应消息：%+v`, charge), "")
	//-------------------模拟点击：A/P INV-------------------
	tools.Logger(param.Serial, "模拟点击：”A/P INV“按钮", "", "")
	exec, err = request_http.AihHouseExec(order.Login, chargesRows)
	if err != nil {
		return types, err
	}
	if exec.Obj.Ask != "CREATE A/P INVOICE?" {
		tools.Logger(param.Serial, "点击A/P INV按钮错误，请人工核实", fmt.Sprintf(`系统提示：%s`, exec.Obj.Ask), tools.LOG_LEVEL_ERROR)
		return types, nil
	}
	tools.Logger(param.Serial, "点击A/P INV按钮成功，模拟点击“OK”按钮", "", "")
	exec, err = request_http.AihHouseExec2(order.Login, chargesRows)
	if err != nil {
		return types, err
	}
	tools.Logger(param.Serial, "点击“OK”按钮完成", fmt.Sprintf(`系统响应消息：%+v`, exec), "")
	types = 2
	return types, nil
}
