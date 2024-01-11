package ccl_ser

import (
	"fmt"
	"mail_download/model"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/service/ccl_ser/air_transport"
	"mail_download/service/ccl_ser/air_transport/request_http"
	"mail_download/service/ccl_ser/shipping"
	"mail_download/service/excel_ser"
	"mail_download/tools"
	customErr "mail_download/tools/error"
	"strings"
	"time"
)

// CLL 日志记录：邮件只需要发送最终的结果即可，
func CLL(param request_model.CCLParam) error {
	var (
		err     error
		login   response_model.LoginResp
		fileMap = map[string]response_model.ReadPdf{}
	)
	if param.Email != "" && !tools.ValidEmail(param.Email) {
		return customErr.New(customErr.EMAIL_FORMAT_ERROR, "")
	}
	_, ok := tools.ProcessMap.Load(param.Account)
	if ok {
		return customErr.New(customErr.ACCOUNT_TWO_PROCESSES_ERROR, "")
	}
	tools.ProcessMap.Store(param.Account, param.Serial) // 添加本次操作表示
	tools.Logger(param.Serial, fmt.Sprintf(`开始准备工作`), "", "")
	tools.Logger(param.Serial, fmt.Sprintf(`客户端发出CCL操作流程指令，开始执行指令；请求信息：%+v`, param), "", "")
	fileMap, err = ReadPdfGroup(param.Url) //获取指定目录下的所有PDF文件
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		return err
	}
	if len(fileMap) > 500 {
		tools.ProcessMap.Delete(param.Account)
		return customErr.New(customErr.TOO_MANY_FILES, "")
	}
	tools.Logger(param.Serial, "开始模拟登录EFC系统", "", "")
	login, err = Login(param.Account, param.Password) //模拟登录系统，检测账号密码是否正确
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		return err
	}
	if login.Message != "系统验证处理中，请稍侯......." {
		tools.ProcessMap.Delete(param.Account)
		return customErr.New(customErr.CLL_LOGIN_ERROR, "")
	}
	tools.Logger(param.Serial, "模拟登录EFC系统成功", fmt.Sprintf(`登录响应结果：%+v`, login), "")
	go Operate(fileMap, login, param)
	return nil
}

func Operate(fileMap map[string]response_model.ReadPdf, login response_model.LoginResp, param request_model.CCLParam) error {
	var (
		err        error
		log        string
		fileNumber = 1
		queryInfo  response_model.AirQueryInfo
		types      int64
		typesMap   = map[int64]int64{}
	)
	defer func() {
		if r := recover(); r != nil {
			tools.ProcessMap.Delete(param.Account)
			SendMessage(typesMap, param, r.(error))
		}
	}()
	time.Sleep(2 * time.Second) // 先休眠两秒钟，因为CCL系统处理比较慢，等待两秒缓冲时间
	tools.Logger(param.Serial, "开始模拟点击Query MAWB/Query Master功能", "", "")
	queryInfo, err = request_http.GetQueryInfo(login) // 模拟查询信息
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		SendMessage(typesMap, param, err)
		return err
	}
	tools.Logger(param.Serial, "模拟点击Query MAWB/Query Master功能成功", fmt.Sprintf(`调用GetQueryInfo查询表格所需要使用到的参数数据，查询结果：%+v`, queryInfo), "")
	for file, pdf := range fileMap {
		tools.Excel(param.Serial, model.XlsxData{})
		log = fmt.Sprintf(`--------开始操作第%d个文件，共%d个，文件名：%s`, fileNumber, len(fileMap), file)
		tools.Logger(param.Serial, log, "", "")
		log = fmt.Sprintf(`在当前文件中检测出 Amount：%v，Invoice No：%s，主单号（Master No/MAWB No）：%v`, pdf.Amount, pdf.InvoiceNo, pdf.MasterMawbNo)
		tools.Logger(param.Serial, log, "", "")
		for _, no := range pdf.MasterMawbNo {
			order := model.OrderParam{FileName: file, MasterMawbNo: no, QueryInfo: queryInfo, PdfData: pdf, Login: login}
			switch {
			case strings.Index(no, "-") < 0:
				//==========海运单号===========
				types, err = shipping.Shipping(order, param)
			case strings.Index(no, "-") >= 0:
				//==========空运单号===========
				types, err = air_transport.AirTransport(order, param)
			default:
				tools.Logger(param.Serial, fmt.Sprintf(`%s未检测出是海运单号还是空运单号，请人工核实`, file), "", tools.LOG_LEVEL_ERROR)
			}
			if err != nil {
				tools.ProcessMap.Delete(param.Account)
				SendMessage(map[int64]int64{}, param, err)
				return err
			}
			count, ok := typesMap[types]
			if ok {
				count++
			}
			typesMap[types] = count
			if types == 2 {
				break //如果这个单号去操作流程并且成功了，这个文件中后面匹配到的单号就没必要再去查询了
			}
		}
		fileNumber++
	}
	SendMessage(typesMap, param, nil)
	tools.ProcessMap.Delete(param.Account)
	return nil
}

func SendMessage(typesMap map[int64]int64, param request_model.CCLParam, r error) {
	var (
		result string
		err    error
		text   = "后续若发现操作的数据存在错误等信息，也可查看记录文件排查问题原因"
	)
	switch {
	case len(typesMap) == 2: //说明有成功和失败的
		if typesMap[1] >= typesMap[2] {
			result = fmt.Sprintf(`CCL流程操作已完成；此次操作仅有部分成功，其余大部分操作因数据原因未执行完成（或被终止），打开记录文件查看详细信息；%s`, text)
		} else if typesMap[1] < typesMap[2] {
			result = fmt.Sprintf(`CCL流程操作已完成；部分操作因数据原因未执行完成（或被终止），打开记录文件查看详细信息；%s`, text)
		}
	case len(typesMap) == 1:
		_, ok1 := typesMap[1]
		_, ok2 := typesMap[2]
		if ok1 {
			result = "CCL流程操作因数据问题导致提供的PDF文件操作全部失败，打开记录文件查看详细信息；" + text
		} else if ok2 {
			result = "CCL流程操作已全部操作成功；" + text
		}
	default:
		result = "CCL流程操作因服务器或CCL系统原因导致中断，请联系相关技术人员处理（已执行完成的主单号不受影响）；" + text
	}
	tools.Logger(param.Serial, result, "", "")
	if r != nil {
		tools.Logger(param.Serial, "程序发生错误，请联系技术人员处理", fmt.Sprintf(`错误信息：%s`, r.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
	}
	err = excel_ser.Excel(param.Serial)
	if err != nil {
		tools.Logger(param.Serial, "保存xlsx出现错误", "错误信息："+err.Error(), tools.LOG_LEVEL_SYSTEM_ERROR)
	}
	if param.Email == "" {
		tools.Logger(param.Serial, "检测到用户未选择邮箱通知，本次结果将不会推送", "", "")
	} else {
		tools.Logger(param.Serial, "发送邮件通知，发送账号："+param.Email, "", "")
		err = tools.MailAttachment(param.Email, result, param.Serial)
		if err != nil {
			tools.Logger(param.Serial, "发送邮件通知失败，请联系技术人员处理", fmt.Sprintf(`错误信息：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
		} else {
			tools.Logger(param.Serial, "邮箱通知已发送成功，请留意你的邮箱；如果没有找到，它可能在你的垃圾邮箱中；", "", "")
		}
	}
	tools.XlsxMap.Delete(param.Serial)
}
