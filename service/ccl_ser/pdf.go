package ccl_ser

import (
	"github.com/ledongthuc/pdf"
	"io/fs"
	"io/ioutil"
	"mail_download/model/response_model"
	customErr "mail_download/tools/error"
	"os"
	"regexp"
	"strings"
)

func ReadPdfGroup(path string) (map[string]response_model.ReadPdf, error) {
	var (
		files     []fs.FileInfo
		err       error
		dataMap   = map[string]response_model.ReadPdf{}
		invoiceNo string
		amounts   string
	)
	files, err = ioutil.ReadDir(path)
	if err != nil {
		return dataMap, customErr.New(customErr.READING_FILE_ERROR, "")
	}
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), "pdf") {
			return dataMap, customErr.New(customErr.DIRECTORY_FILE_ERROR, "")
		}
	}
	for _, file := range files {
		var (
			f            *os.File
			r            *pdf.Reader
			data         = response_model.ReadPdf{}
			masterMawbNo = map[string]struct{}{}
		)
		f, r, err = pdf.Open(path + "\\" + file.Name())
		defer func() {
			_ = f.Close()
		}()
		if err != nil {
			return dataMap, customErr.New(customErr.OPEN_FILE_ERROR, "")
		}
		for pageIndex := 1; pageIndex <= r.NumPage(); pageIndex++ {
			p := r.Page(pageIndex)
			if p.V.IsNull() {
				continue
			}
			rows, _ := p.GetTextByRow()
			for _, row := range rows {
				for _, word := range row.Content {
					master := regexp.MustCompile(`^[A-Z]{4,}.*\d.*$`).MatchString(word.S)
					mawb := regexp.MustCompile(`^[0-9]{3}-[0-9]+$`).MatchString(word.S)
					invoice := regexp.MustCompile(`^CCL-AR\d+$`).MatchString(word.S)
					amount := regexp.MustCompile(`^\d{1,3}(,\d{3})*(\.\d{2})?$`).MatchString(word.S)
					//fmt.Println(word.S)
					if master && !containsInvalidChars(word.S, " /abcdefghijklmnopqrstuvwxyz") {
						masterMawbNo[word.S] = struct{}{}
					}
					if mawb && isValidHyphen(word.S) {
						masterMawbNo[word.S] = struct{}{}
					}
					if invoice {
						invoiceNo = word.S
					}
					if amount {
						amounts = word.S
					}
				}
			}
		}
		for k := range masterMawbNo {
			data.MasterMawbNo = append(data.MasterMawbNo, k)
		}
		amounts = strings.ReplaceAll(amounts, ",", "")
		data.Amount = strings.ReplaceAll(amounts, "，", "")
		data.InvoiceNo = invoiceNo

		filename := file.Name()
		filename = strings.ReplaceAll(filename, " ", " ")
		dataMap[filename] = data
	}
	return dataMap, nil
}

func isValidHyphen(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			count++
			if count > 1 || i != 3 {
				return false
			}
		} else if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return count == 1
}

func containsInvalidChars(s string, invalidChars string) bool {
	for _, char := range invalidChars {
		if containsChar(s, byte(char)) {
			return true
		}
	}
	return false
}

func containsChar(s string, char byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == char {
			return true
		}
	}
	return false
}
