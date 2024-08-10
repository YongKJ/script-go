package GenUtil

import (
	"bufio"
	"container/list"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	lzstring "github.com/daku10/go-lz-string"
	"github.com/golang-module/carbon/v2"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/FileUtil"
	"script-go/src/application/util/LogUtil"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func Print(str ...string) {
	fmt.Print(strings.Join(str, ""))
}

func Println(str ...string) {
	fmt.Println(strings.Join(str, ""))
}

func ReadParams() []string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	return StrToArray(strings.TrimSpace(text), " ")
}

func GetMapKeys(mapData map[string]any) []string {
	var lstKey []string
	for key := range mapData {
		lstKey = append(lstKey, key)
	}
	return lstKey
}

func GetValue(key string) any {
	return GetConfig(GetYaml())[key]
}

func GetConfig(config string) map[string]any {
	path := GetConfigPath(config)
	content := FileUtil.Read(path)

	mapData := make(map[string]any)
	err := yaml.Unmarshal([]byte(content), &mapData)
	if err != nil {
		log.Println(err)
	}

	return mapData
}

func GetConfigPath(config string) string {
	execPath, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	execDir := filepath.Dir(execPath)
	path := filepath.Join(execDir, config)
	if FileUtil.Exist(path) {
		return path
	}
	path = FileUtil.GetAbsPath(config)
	if FileUtil.Exist(path) {
		return path
	}
	return FileUtil.GetAbsPath("src", "assets", config)
}

func GetYaml() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	if strings.Contains(execPath, "script_go") ||
		strings.Contains(execPath, "script_go_test") {
		return GetYamlByContent(execPath)
	}
	index := strings.LastIndex(execPath, string(filepath.Separator))
	return ToLine(strings.Replace(execPath[index+1:], ".exe", "", 1) + ".yaml")
}

func GetYamlByContent(execPath string) string {
	appName := "Application.go"
	if strings.Contains(execPath, "script_go_test") {
		appName = "ApplicationTest.go"
	}
	appPath := FileUtil.GetAbsPath("src", "application", appName)
	lines := FileUtil.ReadByLine(appPath)
	regex := regexp.MustCompile("\\s+(\\S+)\\.Run\\(\\)")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "//") {
			continue
		}
		match := regex.MatchString(line)
		if !match {
			continue
		}
		parts := regex.FindStringSubmatch(line)
		if len(parts) == 0 {
			continue
		}
		return ToLine(parts[1]) + ".yaml"
	}
	return ""
}

func ArraysToAnyByRef(arrays any) []any {
	values := reflect.ValueOf(arrays)
	lstData := make([]any, values.Len())
	for i := 0; i < values.Len(); i++ {
		lstData[i] = values.Index(i).Interface()
	}
	return lstData
}

func ArraysToAny[A ~[]E, E any](arrays A) []any {
	lstData := make([]any, len(arrays))
	for i, array := range arrays {
		lstData[i] = array
	}
	return lstData
}

func GetDecodeOSPaths(paths []string) []string {
	lstPath := make([]string, len(paths))
	for i, path := range paths {
		lstPath[i] = GetDecodeOSPath(path)
	}
	return lstPath
}

func GetDecodeOSPathBySep(path string) []string {
	return GetDecodeOSPaths(ParamStrToArray(path, true))
}

func GetDecodeOSPath(path string) string {
	return GetPathByOS(GetDecode(path))
}

func ParamStrToArray(path string, isDecode bool) []string {
	var paths []string
	if strings.Contains(path, "/") {
		paths = StrToArray(path, "/")
	}
	if strings.Contains(path, ",") {
		paths = StrToArray(path, ",")
	}
	return ParamStrToArrayByDecode(paths, isDecode)
}

func ParamStrToArrayByDecode(paths []string, isDecode bool) []string {
	if isDecode {
		return GetDecodes(paths)
	}
	return paths
}

func GetDecodes(lstStr []string) []string {
	strs := make([]string, len(lstStr))
	for i, str := range lstStr {
		strs[i] = GetDecode(str)
	}
	return strs
}

func GetDecode(str string) string {
	return GetDecodeStr(str)
}

func GetEncodes(lstStr []string) []string {
	strs := make([]string, len(lstStr))
	for i, str := range lstStr {
		strs[i] = GetEncode(str)
	}
	return strs
}

func GetEncode(str string) string {
	return GetEncodeStr(str)
}

func GetDecodeStr(str string) string {
	decompressedStr := decodeURIsafe(str)
	message, err := base64.StdEncoding.DecodeString(decompressedStr)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("GenUtil", "GetDecodeStr", "base64.StdEncoding.DecodeString", err))
	}
	return string(message)
}

func GetEncodeStr(str string) string {
	strBase64 := base64.StdEncoding.EncodeToString([]byte(str))
	return encodeWithURIsafe(strBase64)
}

func encodeWithURIsafe(content string) string {
	contentBase64 := base64.StdEncoding.EncodeToString([]byte(content))
	compressedStr, err := lzstring.CompressToEncodedURIComponent(contentBase64)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("GenUtil", "ApplicationTest", "lzstring.CompressToEncodedURIComponent", err))
	}
	return compressedStr
}

func decodeURIsafe(content string) string {
	decompressedStr, err := lzstring.DecompressFromEncodedURIComponent(content)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("GenUtil", "decodeURIsafe", "lzstring.DecompressFromEncodedURIComponent", err))
	}
	message, err := base64.StdEncoding.DecodeString(decompressedStr)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("GenUtil", "decodeURIsafe", "base64.StdEncoding.DecodeString", err))
	}
	return string(message)
}

func GetPathsByOS(paths []string) []string {
	if runtime.GOOS == "windows" {
		return paths
	}
	lstPath := make([]string, len(paths))
	for i, path := range lstPath {
		lstPath[i] = GetPathByOS(path)
	}
	return lstPath
}

func GetPathByOS(path string) string {
	if runtime.GOOS == "windows" {
		return path
	}
	index := strings.Index(path, string(filepath.Separator))
	if index == -1 && len(path) > 0 {
		return "/"
	}
	return path[index:]
}

func GetMd5Str(str string) string {
	hash := md5.Sum([]byte(str))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

func IsEmpty(object interface{}) bool {
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		} else {
			return false
		}
	}
	return false
}

func StrToArray(str string, separator string) []string {
	if len(str) == 0 {
		return []string{}
	}

	if len(separator) == 0 {
		separator = " "
	}

	str = strings.TrimSpace(str)
	return strings.Split(str, separator)
}

func ArrayToStr(lstStr []string, separator string) string {
	tempStr := ""
	for _, str := range lstStr {
		if len(separator) == 0 {
			tempStr += str + " "
			continue
		}
		tempStr += str + separator
	}
	if len(separator) > 0 {
		return tempStr[0 : len(tempStr)-len(separator)]
	}
	return tempStr[0 : len(tempStr)-1]
}

func StrToJsonArray(jsonStr string) []any {
	var lstData []any
	err := json.Unmarshal([]byte(jsonStr), &lstData)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("GenUtil", "StrToJsonArray", "err", err))
	}
	return lstData
}

func StrToMap(jsonStr string) map[string]any {
	var mapData map[string]any
	err := json.Unmarshal([]byte(jsonStr), &mapData)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("GenUtil", "StrToMap", "err", err))
	}
	return mapData
}

func ArrayToList(arrayData []any) *list.List {
	lstData := list.New()
	for i := 0; i < len(arrayData); i++ {
		lstData.PushBack(arrayData[i])
	}
	return lstData
}

func ListToStrArray(lstData *list.List) []string {
	i := 0
	arrayData := make([]string, lstData.Len())
	for data := lstData.Front(); data != nil; data = data.Next() {
		arrayData[i] = data.Value.(string)
		i++
	}
	return arrayData
}

func ListToArray(lstData *list.List) []any {
	i := 0
	arrayData := make([]any, lstData.Len())
	for data := lstData.Front(); data != nil; data = data.Next() {
		arrayData[i] = data.Value
		i++
	}
	return arrayData
}

func ArrayRegContain(array []string, element string) bool {
	for _, regStr := range array {
		regex := regexp.MustCompile(regStr)
		match := regex.MatchString(element)
		if match {
			return true
		}
	}
	return false
}

func ArrayContain(array []string, element string) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

func ArrayContains(array []string, element string) bool {
	for _, value := range array {
		if strings.Contains(element, value) {
			return true
		}
	}
	return false
}

func DateToStr(dateTime time.Time) string {
	return carbon.CreateFromStdTime(dateTime).Format("Y-m-d H:i:s")
}

func StrToDate(dateTimeStr string) time.Time {
	return carbon.Parse(dateTimeStr).StdTime()
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func ToHump(name string) string {
	regex := regexp.MustCompile("\\-(\\w)")
	return strings.ToUpper(name[0:1]) + regex.ReplaceAllStringFunc(name[1:], func(str string) string {
		return strings.ToUpper(strings.Replace(str, "-", "", 1))
	})
}

func ToLine(name string) string {
	regex := regexp.MustCompile("([A-Z])")
	return strings.ToLower(name[0:1]) + regex.ReplaceAllStringFunc(name[1:], func(str string) string {
		return "-" + strings.ToLower(str)
	})
}

func GetSizeStr(size float64) string {
	if size == -1 {
		return ""
	}
	tempSize := size / 1024 / 1024 / 1024 / 1024 / 1024 / 1024
	sizeStr := Float64ToString(Decimal(tempSize)) + " EB"
	if tempSize < 1 {
		tempSize = size / 1024 / 1024 / 1024 / 1024 / 1024
		sizeStr = Float64ToString(Decimal(tempSize)) + " PB"
	}
	if tempSize < 1 {
		tempSize = size / 1024 / 1024 / 1024 / 1024
		sizeStr = Float64ToString(Decimal(tempSize)) + " TB"
	}
	if tempSize < 1 {
		tempSize = size / 1024 / 1024 / 1024
		sizeStr = Float64ToString(Decimal(tempSize)) + " GB"
	}
	if tempSize < 1 {
		tempSize = size / 1024 / 1024
		sizeStr = Float64ToString(Decimal(tempSize)) + " MB"
	}
	if tempSize < 1 {
		tempSize = size / 1024
		sizeStr = Float64ToString(Decimal(tempSize)) + " KB"
	}
	if tempSize < 1 {
		sizeStr = Float64ToString(Decimal(size)) + " B"
	}
	return sizeStr
}

func Decimal(num float64) float64 {
	num, err := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	if err != nil {
		log.Println(err)
	}
	return num
}

func StrToInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Println(err)
	}
	return num
}

func StrToBoolean(str string) bool {
	return str == "true"
}

func StrToFloat64(numStr string) float64 {
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		log.Println(err)
	}
	return num
}

func BoolToString(boolData bool) string {
	return strconv.FormatBool(boolData)
}

func MapToString(mapData map[string]any) string {
	jsonStr, err := json.Marshal(mapData)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(jsonStr)
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Int64ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}

func Uint64ToString(number uint64) string {
	return strconv.FormatUint(number, 10)
}

func Float64ToString(number float64) string {
	return strconv.FormatFloat(number, 'f', -1, 64)
}
