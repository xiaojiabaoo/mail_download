package tools

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func GetUuid() string {
	return uuid.NewV4().String()
}

func GetUuidStr() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func RandNum(min, max int64) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	code := r.Int63n(max-min) + min
	return Int2Str(int(code))
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntToBool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	value, _ := strconv.Atoi(s)
	return value
}

func StringToInt32(s string) int32 {
	if s == "" {
		return 0
	}
	value, _ := strconv.Atoi(s)
	return int32(value)
}

func Float64ToStr(float float64) string {
	return strconv.FormatFloat(float, 'f', -1, 64)
}

func StringToInt64(s string) int64 {
	value, _ := strconv.ParseInt(s, 10, 64)
	return value
}

func StringToFloat64(s string) float64 {
	value, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return 0.0
	}
	return value
}

func AnyToString(data any) string {
	if data != nil {
		return data.(string)
	}
	return ""
}

// 字符串转int,失败返回默认值,没有默认值返回）
func Int2Str(i int) string {
	return strconv.Itoa(i)
}

func Str2Int(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		res = 0
	}
	return res
}

func Str2Int64(s string) int64 {
	if s == "" || len(s) == 0 {
		return 0
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		res = 0
	}
	return int64(res)
}

// 查看字符串是否包含子字符串
func HasString(param ...string) bool {
	if len(param) < 2 {
		return false
	}

	str := param[0]

	for i := 1; i < len(param); i++ {
		if strings.Contains(str, param[i]) {
			return true
		}
	}
	return false
}

func SliceIntMap(str, sub string, mapData map[string]int) []int {
	params := strings.Split(str, sub)
	data := make([]int, 0)

	for _, param := range params {
		if param != "" {
			num, ok := mapData[param]
			if ok {
				data = append(data, num)
			}
		}

	}
	return data
}

func SliceString2Int(str, sub string) []int {
	params := strings.Split(str, sub)
	data := make([]int, 0)

	for _, param := range params {
		if param != "" {
			num, err := strconv.Atoi(param)
			if err == nil {
				data = append(data, num)
			}
		}

	}
	return data
}

func SliceStringMap(str, sub string, mapData map[string]string) []string {
	data := make([]string, 0)
	if mapData == nil {
		return data
	}

	params := strings.Split(str, sub)
	for _, param := range params {
		if param != "" {
			value, ok := mapData[param]
			if ok {
				data = append(data, value)
			}
		}

	}
	return data
}

func SliceString(str, sub string) []string {
	params := strings.Split(str, sub)
	data := make([]string, 0)

	for _, param := range params {
		if param != "" {
			data = append(data, param)
		}
	}
	return data
}

func SliceInt(str, sub string) []int {
	params := strings.Split(str, sub)
	data := make([]int, 0)
	for _, param := range params {
		if param != "" {
			data = append(data, Str2Int(param))
		}
	}
	return data
}

func SliceToString(data []string, sub string) string {
	if len(data) == 0 {
		return ""
	}
	str := ""
	for k, v := range data {
		if k != len(data)-1 {
			str += v + sub
		} else {
			str += v
		}

	}
	return str
}

// 切片 变成字符串 默认,号连接
func SliceIntToString(s []string, old, new string) string {
	if new == "" {
		new = ","
	}
	if old == "" {
		old = " "
	}
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), old, new, -1)
}

// 切片变执行sqls时的in条件
func SliceStringToString(s []string, old, new string) string {
	if new == "" {
		new = "','"
	}
	if old == "" {
		old = " "
	}
	return "'" + strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), old, new, -1) + "'"
}

// 整数转浮点型百分比
func IntAccuracy(i, j int) float64 {
	if i == 0 || j == 0 {
		return 0
	}
	return float64(i) / float64(j)
}

func FloatString(num float64, base int) string {
	if base > 4 {
		base = 4
	}
	if base == 0 {
		base = 1
	}

	str := ""
	switch base {
	case 0:
		{
			str = fmt.Sprintf("%0.0f", num)
		}
	case 1:
		{
			str = fmt.Sprintf("%0.1f", num)
		}
	case 2:
		{
			str = fmt.Sprintf("%0.2f", num)
		}
	case 3:
		{
			str = fmt.Sprintf("%0.3f", num)
		}
	case 4:
		{
			str = fmt.Sprintf("%0.4f", num)
		}
	}
	return str
}

// 整数转浮点型百分比字符串 最大支持小数点后四位
func IntAccuracyString(i, j, base int) string {
	if base > 4 {
		base = 4
	}
	if i == 0 || j == 0 {
		return "0"
	}

	num := float64(i) / float64(j)
	str := ""
	switch base {
	case 0:
		{
			str = fmt.Sprintf("%0.0f", num)
		}
	case 1:
		{
			str = fmt.Sprintf("%0.1f", num)
		}
	case 2:
		{
			str = fmt.Sprintf("%0.2f", num)
		}
	case 3:
		{
			str = fmt.Sprintf("%0.3f", num)
		}
	case 4:
		{
			str = fmt.Sprintf("%0.4f", num)
		}
	}
	return str
}

// 从计数的map中取出排在前面的count个
func GetMapSortN(mapData map[string]int, count int) []string {
	if mapData == nil {
		return nil
	}

	mapTemp := make(map[int][]string)
	sortTemp := make([]int, 0)
	for k, v := range mapData {
		temp, ok := mapTemp[v]
		if ok {
			temp = append(temp, k)
			mapTemp[v] = temp
		} else {
			sortTemp = append(sortTemp, v)
			mapTemp[v] = []string{k}
		}
	}
	sort.Ints(sortTemp)
	num := count
	data := make([]string, 0)
	for i := len(sortTemp) - 1; i >= 0; i-- {
		sli := mapTemp[sortTemp[i]]
		for _, v := range sli {
			num--
			data = append(data, v)
			if num == 0 {
				return data
			}
		}
	}
	return data
}

// 去除字符串中最后一个字符
func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

// 去除字符串的空格  换行 制表符
func StringFmt(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func MapFmt(data map[string]interface{}) {
	if data == nil {
		return
	}

	for k, v := range data {
		value, ok := v.(string)
		if ok {
			if value == "" {
				delete(data, k)
			}
		}
	}
}

func CreateUuid() string {
	return uuid.NewV4().String()
}

func ValidEmail(email string) bool {
	// 定义邮箱地址的正则表达式
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 使用正则表达式匹配邮箱地址
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

// 并发安全对象
var snowFlakeID *snowflake.Node

func CreateSerialNumber() string {
	if snowFlakeID == nil {
		snowFlakeID, _ = snowflake.NewNode(1)
	}

	id := snowFlakeID.Generate()
	return strconv.FormatInt(int64(id), 10)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 公共get请求方法带失败重试
func Send(url string, try int) ([]byte, error) {
	result := make([]byte, 0)
	var err error
	for i := 0; i < try; i++ {
		result, err = httpReq(url)
		if err == nil {
			break
		}
	}
	return result, err
}

// 发起http请求
func httpReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// 3、发送http请求
	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetPageInfo(page, pageSize int) (first int, pageItemNumber int) {
	if page <= 0 || page > 1000000 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 20
	}
	pageItemNumber = pageSize
	first = (page - 1) * pageItemNumber
	return
}

// FlateEncode 数据压缩
// levels range from 1 (BestSpeed) to 9 (BestCompression)
func FlateEncode(input []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w, err := flate.NewWriter(buf, 9)
	if err != nil {
		return nil, err
	}
	w.Write(input)
	w.Close()
	return buf.Bytes(), nil
}

// FlateDecode 数据解压
func FlateDecode(data []byte) ([]byte, error) {
	return ioutil.ReadAll(flate.NewReader(bytes.NewBuffer(data)))
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}

// IsEnglishText 检测字符串是否以字母开头
func IsEnglishText(str string) bool {
	u := str[0]
	if (u >= 65 && u <= 90) || (u >= 97 && u <= 122) {
		return true
	}
	return false
}

func Md5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// 检测字符串是否存在中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d" +
			"\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

func MyZip(fileName string, zipFileName string) error {
	// 创建：zip文件
	zipfile, _ := os.Create(zipFileName)
	defer zipfile.Close()
	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	// 遍历路径信息

	info, err := os.Lstat(fileName)
	if err != nil {
		return err
	}
	// 获取：文件头信息
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = info.Name()

	// 判断：文件是不是文件夹

	// 设置：zip的文件压缩算法
	header.Method = zip.Deflate

	// 创建：压缩包头部信息
	writer, _ := archive.CreateHeader(header)
	file, _ := os.Open(fileName)
	defer file.Close()
	_, err = io.Copy(writer, file)
	return err
}

// 压缩为zip文件
func Compress(file string, dest string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()

	wr := zip.NewWriter(d)
	defer wr.Close()

	w, err := wr.Create(file)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, f)
	if err != nil {
		return err
	}

	return nil
}

// 距离现在几年
func GetFromNowByYear(year int) (age int) {
	if year <= 0 {
		age = 0
	}
	nowYear := time.Now().Year()
	age = nowYear - year
	return
}

// TruncateDecimal 四舍五入小数操作
// value 需要操作的小数
// number 四舍五入后保留几位小数
func TruncateDecimal(value float64, number int32) float64 {
	if value == 0 {
		return value
	}
	data, _ := decimal.NewFromFloat(value).Round(number).Float64()
	return data
}

func GetAge(year int) (age int) {
	if year <= 0 {
		age = 0
	}
	nowYear := time.Now().Year()
	age = nowYear - year
	return
}

// DivisionProbTwoDecimal 除法处理，结果保留两位小数
func DivisionProbTwoDecimal(divisor, dividend float64) float64 {
	if dividend == 0 {
		return 0
	}
	result, _ := decimal.NewFromFloat(divisor / dividend).Round(2).Float64()
	return result
}

// DivisionProbAnyDecimal 除法处理，结果保留指定位数的小数
func DivisionProbAnyDecimal(divisor, dividend float64, number int32) float64 {
	if dividend == 0 {
		return 0
	}
	result, _ := decimal.NewFromFloat(divisor / dividend).Round(number).Float64()
	return result
}

// DivisionProbAllDecimal 除法处理，结果保留全部小数
func DivisionProbAllDecimal(divisor, dividend float64) float64 {
	if dividend == 0 {
		return 0
	}
	return divisor / dividend
}

func PassMd5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func Serial(serial string) string {
	serial = strings.ReplaceAll(serial, "MAIL-", "")
	serial = strings.ReplaceAll(serial, "CCL-", "")
	return serial
}
