package excel_ser

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/pkg/errors"
	"mail_download/model"
	"mail_download/tools"
)

func Excel(serial string) error {
	var (
		file  *excelize.File
		err   error
		title = &[]string{"操作时间", "记录等级", "记录描述", "操作响应数据"}
		path  string
		data  = make([]model.XlsxData, 0)
	)
	value, ok := tools.XlsxMap.Load(serial)
	if !ok {
		return errors.New("获取xlsx的记录数据为空，xlsx终止保存")
	}
	data = value.([]model.XlsxData)
	path = fmt.Sprintf(`./logs/%s/%s.xlsx`, tools.Serial(serial), serial)
	// 判断文件是否存在
	file, err = excelize.OpenFile(path)
	if err != nil {
		file = excelize.NewFile()
		file.SetActiveSheet(0)                  // 设置默认工作表
		file.SetSheetRow("Sheet1", "A1", title) // 设置title表头
		file.SetRowHeight("Sheet1", 1, 18)      // 设置表头的高度
		style, _ := file.NewStyle(`{"fill":{"type":"pattern","color":["#D6EAF8"],"pattern":1},"alignment":{"horizontal":"center"}}`)
		file.SetCellStyle("Sheet1", "A1", "D1", style) // 设置表头的背景
		// 设置表头首行首行冻结
		file.SetPanes("Sheet1", `{"freeze":true,"split":false,"x_split":0,"y_split":1,"top_left_cell":"A1","active_pane":"bottomLeft","panes":[{"sq_ref":"A1","active_cell":"D1"}]}`)
		//file.SetColWidth("Sheet1", "A", "D", 10) // 设置列宽 A ~ D
		file.SetColWidth("Sheet1", "A", "A", 20)  // 设置列宽 A ~ D
		file.SetColWidth("Sheet1", "B", "B", 10)  // 设置列宽 A ~ D
		file.SetColWidth("Sheet1", "C", "C", 100) // 设置列宽 A ~ D
		file.SetColWidth("Sheet1", "D", "D", 50)  // 设置列宽 A ~ D
	}
	emptyRow := file.GetRows("Sheet1")
	for k, v := range data {
		var (
			lint       = k + len(emptyRow) + 1
			level      string
			background int
		)
		if v.Level == "" {
			file.MergeCell("Sheet1", fmt.Sprintf("A%d", lint), fmt.Sprintf("D%d", lint))
			background, _ = file.NewStyle(`{"fill":{"type":"pattern","color":["#EAEDED"],"pattern":1}}`)
			file.SetCellStyle("Sheet1", fmt.Sprintf("A%d", lint), fmt.Sprintf("D%d", lint), background)
			continue
		}
		switch v.Level {
		case "INFO":
			level = "正常"
			background, _ = file.NewStyle(`{"fill":{"type":"pattern","color":["#D5F5E3"],"pattern":1},"alignment":{"horizontal":"center"}}`)
		case "ERROR":
			level = "逻辑错误"
			background, _ = file.NewStyle(`{"fill":{"type":"pattern","color":["#FCF3CF"],"pattern":1},"alignment":{"horizontal":"center"}}`)
		case "SYSTEM_ERROR":
			level = "系统错误"
			background, _ = file.NewStyle(`{"fill":{"type":"pattern","color":["#FADBD8"],"pattern":1},"alignment":{"horizontal":"center"}}`)
		}
		// 设置”记录等级“列的背景颜色和居中显示
		file.SetCellStyle("Sheet1", fmt.Sprintf("B%d", lint), fmt.Sprintf("B%d", lint), background)
		// 设置”操作时间“列的居中显示
		background, _ = file.NewStyle(`{"alignment":{"horizontal":"center"}}`)
		file.SetCellStyle("Sheet1", fmt.Sprintf("A%d", lint), fmt.Sprintf("A%d", lint), background)
		// 写入数据
		file.SetSheetRow("Sheet1", fmt.Sprintf("A%d", lint), &[]interface{}{
			v.Time,     // 操作时间
			level,      // 记录等级
			v.Describe, // 记录描述
			v.Response, // 操作第三方响应数据
		})
		file.SetRowHeight("Sheet1", lint, 18)
	}
	err = file.SaveAs(path)
	if err != nil {
		return errors.Wrap(err, "保存文件出错")
	}
	return nil
}
