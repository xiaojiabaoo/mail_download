package ccl_ser

import (
	"fmt"
	"mail_download/model"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/service/ccl_ser/air_transport"
	"mail_download/service/ccl_ser/air_transport/request_http"
	"mail_download/service/ccl_ser/shipping"
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
	tools.Logger(param.Serial, fmt.Sprintf(`-----------------------分隔符-开始准备工作------------------------`), "")
	tools.Logger(param.Serial, fmt.Sprintf(`客户端发出CCL操作流程指令，开始执行指令；请求信息：%+v`, param), "")
	fileMap, err = ReadPdfGroup(param.Url) //获取指定目录下的所有PDF文件
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		return err
	}
	if len(fileMap) > 500 {
		tools.ProcessMap.Delete(param.Account)
		return customErr.New(customErr.TOO_MANY_FILES, "")
	}
	tools.Logger(param.Serial, "开始模拟登录CCL系统", "")
	login, err = Login(param.Account, param.Password) //模拟登录系统，检测账号密码是否正确
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		return err
	}
	if login.Message != "系统验证处理中，请稍侯......." {
		tools.ProcessMap.Delete(param.Account)
		return customErr.New(customErr.CLL_LOGIN_ERROR, "")
	}
	tools.Logger(param.Serial, fmt.Sprintf(`模拟登录CCL系统成功，系统登录响应结果：%+v`, login), "")
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
			sendMessage(typesMap, param)
		}
	}()
	time.Sleep(2 * time.Second) // 先休眠两秒钟，因为CCL系统处理比较慢，等待两秒缓冲时间
	tools.Logger(param.Serial, "开始模拟点击Query MAWB/Query Master功能", "")
	queryInfo, err = request_http.GetQueryInfo(login) // 模拟查询信息
	if err != nil {
		tools.ProcessMap.Delete(param.Account)
		go tools.MailAttachment(param.Email, tools.CCL_RESULT_FAIL, param.Serial)
		return err
	}
	tools.Logger(param.Serial, fmt.Sprintf(`模拟点击Query MAWB/Query Master功能成功，调用GetQueryInfo查询表格所需要使用到的参数数据，查询结果：%+v`, queryInfo), "")
	for file, pdf := range fileMap {
		log = fmt.Sprintf(`-----------------------分隔符-开始操作第%d个文件，共%d个，文件名：%s------------------------`, fileNumber, len(fileMap), file)
		tools.Logger(param.Serial, log, "")
		log = fmt.Sprintf(`在当前文件中检测出 Amount：%v，Invoice No：%s，主单号（Master No/MAWB No）：%v`, pdf.Amount, pdf.InvoiceNo, pdf.MasterMawbNo)
		tools.Logger(param.Serial, log, "")
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
				tools.Logger(param.Serial, fmt.Sprintf(`%s未检测出是海运单号还是空运单号，请人工核实`, file), tools.LOG_LEVEL_ERROR)
			}
			if err != nil {
				tools.ProcessMap.Delete(param.Account)
				tools.Logger(param.Serial, fmt.Sprintf(`服务器错误导致流程中断，请联系技术人员，错误信息：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
				go tools.MailAttachment(param.Email, tools.CCL_RESULT_FAIL, param.Serial)
				return err
			}
			if types == 2 {
				continue //如果这个单号去操作流程并且成功了，后面匹配到的单号就没必要再去查询了
			}
			count, ok := typesMap[types]
			if ok {
				count++
			}
			typesMap[types] = count
		}
		fileNumber++
	}
	sendMessage(typesMap, param)
	tools.ProcessMap.Delete(param.Account)
	return nil
}

func sendMessage(typesMap map[int64]int64, param request_model.CCLParam) {
	var (
		result string
		err    error
	)
	if param.Email == "" {
		return
	}
	tools.Logger(param.Serial, "准备发送邮件通知", "")
	switch {
	case len(typesMap) == 2: //说明有成功和失败的
		if typesMap[1] >= typesMap[2] {
			result = tools.CCL_RESULT_PART_SUCCESS
		} else if typesMap[1] < typesMap[2] {
			result = tools.CCL_RESULT_PART_FAIL
		} else {
			result = tools.CCL_RESULT_PART_SUCCESS
		}
	case len(typesMap) == 1:
		_, ok1 := typesMap[1]
		_, ok2 := typesMap[2]
		if ok1 {
			result = tools.CCL_RESULT_ALL_FAIL
		} else if ok2 {
			result = tools.CCL_RESULT_ALL_SUCCESS
		}
	}
	err = tools.MailAttachment(param.Email, result, param.Serial)
	if err != nil {
		tools.Logger(param.Serial, fmt.Sprintf(`发送邮件通知失败，请联系技术人员处理；错误信息：%s`, err.Error()), tools.LOG_LEVEL_SYSTEM_ERROR)
	}
	tools.Logger(param.Serial, "邮箱通知已发送成功", "")
}
