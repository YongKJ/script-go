package FileUtil

import (
	"bufio"
	"container/list"
	"github.com/djherbis/times"
	"github.com/integralist/go-findroot/find"
	"github.com/shirou/gopsutil/disk"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/LogUtil"
	"strings"
	"time"
)

var baseDir = ""
var _ = AppDir()

func Disks() []*disk.UsageStat {
	parts, err := disk.Partitions(true)
	disks := make([]*disk.UsageStat, len(parts))
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Disks", "disk.Partitions", err))
		return disks
	}

	i := 0
	for _, part := range parts {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			LogUtil.LoggerLine(Log.Of("FileUtil", "Disks", "disk.Usage", err))
		}
		disks[i] = usage
		i += 1
	}
	return disks
}

func AppDir() string {
	if len(baseDir) > 0 {
		return baseDir
	}

	rootDir, err := find.Repo()
	if err != nil {
		appDir, _ := filepath.Abs(rootDir.Path)
		baseDir = appDir
		if runtime.GOOS == "windows" && strings.Contains(baseDir, "/") {
			baseDir = strings.Replace(baseDir, "/", "\\", -1)
		}
		LogUtil.LoggerLine(Log.Of("FileUtil", "AppDir", "find.Repo", err))
		return baseDir
	}

	baseDir = rootDir.Path
	if runtime.GOOS == "windows" && strings.Contains(baseDir, "/") {
		baseDir = strings.Replace(baseDir, "/", "\\", -1)
	}
	return baseDir
}

func GetAbsPath(names ...string) string {
	return filepath.Join(baseDir, filepath.Join(names...))
}

func Desktop() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	if runtime.GOOS == "windows" {
		return filepath.Join(homeDir, "Desktop")
	}
	return homeDir
}

func WorkFolder() string {
	folder, err := os.Getwd()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "WorkFolder", "os.Getwd", err))
	}
	return folder
}

func Create(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	file.Close()
}

func Size(fileName string) int64 {
	file, err := os.Stat(fileName)
	if err != nil {
		return 0
	}
	return file.Size()
}

func SizeFolder(fileName string) int64 {
	folderSize := int64(0)
	files := List(fileName)
	for i := 0; i < len(files); i++ {
		tempFileName := fileName + string(filepath.Separator) + files[i]
		if IsFolder(tempFileName) {
			folderSize += SizeFolder(tempFileName)
		} else {
			folderSize += Size(tempFileName)
		}
	}
	return folderSize
}

func Exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func Type(fileName string) string {
	if Size(fileName) == 0 {
		return ""
	}
	file, err := os.Open(fileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Type", "os.Open", err))
		return ""
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Type", "file.Read", err))
		return ""
	}

	return http.DetectContentType(buffer[:n])
}

func Date(fileName string) time.Time {
	fileTime, err := times.Stat(fileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Date", "times.Stat", err))
		return time.Now()
	}
	return fileTime.BirthTime()
}

func ModDate(fileName string) time.Time {
	fileTime, err := times.Stat(fileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "ModDate", "times.Stat", err))
		return time.Now()
	}
	return fileTime.ModTime()
}

func IsFolder(fileName string) bool {
	file, err := os.Stat(fileName)
	if err != nil {

		return false
	}
	return file.IsDir()
}

func IsFile(fileName string) bool {
	return !IsFolder(fileName)
}

func Mkdir(fileName string) {
	err := os.MkdirAll(fileName, 0755)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Mkdir", "os.MkdirAll", err))
	}
}

func List(fileName string) []string {
	files, err := os.ReadDir(fileName)
	names := make([]string, len(files))
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "List", "os.ReadDir", err))
		return names
	}
	i := 0
	for _, file := range files {
		names[i] = file.Name()
		i++
	}
	return names
}

func ListFolderByArray(fileName string) []string {
	return listToStrArray(listFolder(fileName))
}

func listToStrArray(lstData *list.List) []string {
	i := 0
	arrayData := make([]string, lstData.Len())
	for data := lstData.Front(); data != nil; data = data.Next() {
		arrayData[i] = data.Value.(string)
		i++
	}
	return arrayData
}

func ListFolder(fileName string) *list.List {
	return listFolder(fileName)
}

func listFolder(fileName string) *list.List {
	lstData := list.New()
	files := List(fileName)
	for i := 0; i < len(files); i++ {
		tempFileName := fileName + string(filepath.Separator) + files[i]
		if IsFolder(tempFileName) {
			lstData.PushBackList(listFolder(tempFileName))
		} else {
			lstData.PushBack(tempFileName)
		}
	}
	return lstData
}

func Read(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Read", "os.ReadFile", err))
	}
	return string(content)
}

func ReadByLine(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "ReadByLine", "os.Open", err))
	}
	defer file.Close()

	lines := list.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines.PushBack(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "ReadByLine", "scanner.Err", err))
	}

	i := 0
	lstLine := make([]string, lines.Len())
	for line := lines.Front(); line != nil; line = line.Next() {
		lstLine[i] = line.Value.(string)
		i++
	}
	return lstLine
}

func ReadByLineAndFunc(fileName string, lineFunc func(line string)) {
	file, err := os.Open(fileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "ReadByLineAndFunc", "os.Open", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineFunc(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "ReadByLineAndFunc", "scanner.Err", err))
	}
}

func WriteStream(fileName string, content []byte) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "WriteStream", "os.OpenFile", err))
		return
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "WriteStream", "file.Write", err))
	}
}

func Write(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Write", "os.OpenFile", err))
		return
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Write", "file.WriteString", err))
	}
}

func Append(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Append", "os.OpenFile", err))
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Append", "file.WriteString", err))
	}
}

func Move(srcFileName string, desFileName string) {
	err := os.Rename(srcFileName, desFileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Move", "os.Rename", err))
	}
}

func Copy(srcFileName string, desFileName string) {
	if IsFolder(srcFileName) {
		Mkdir(desFileName)
		CopyFolder(srcFileName, desFileName)
		return
	}

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Copy", "os.Open", err))
	}
	defer srcFile.Close()

	desFile, err := os.Create(desFileName)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Copy", "os.Create", err))
	}
	defer desFile.Close()

	_, err = io.Copy(desFile, srcFile)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("FileUtil", "Copy", "io.Copy", err))
	}
}

func CopyFolder(srcFolderName string, desFolderName string) {
	files := List(srcFolderName)
	for i := 0; i < len(files); i++ {
		srcNewFileName := srcFolderName + string(filepath.Separator) + files[i]
		desNewFileName := desFolderName + string(filepath.Separator) + files[i]
		if IsFolder(srcNewFileName) {
			Mkdir(desNewFileName)
			CopyFolder(srcNewFileName, desNewFileName)
		} else {
			Copy(srcNewFileName, desNewFileName)
		}
	}
}

func Delete(fileName string) {
	if !Exist(fileName) {
		return
	}

	if IsFile(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			LogUtil.LoggerLine(Log.Of("FileUtil", "Delete", "os.Remove", err))
		}
	}

	if IsFolder(fileName) {
		err := os.RemoveAll(fileName)
		if err != nil {
			LogUtil.LoggerLine(Log.Of("FileUtil", "Delete", "os.RemoveAll", err))
		}
	}
}

func ModFile(path string, regStr string, isAll bool, value string) {
	ModifyFile(path, regStr, isAll, func(matchStr ...string) string {
		return value
	})
}

func ModifyFile(path string, regStr string, isAll bool, valueFunc func(partsStr ...string) string) {
	content := Read(path)
	regex := regexp.MustCompile(regStr)
	if isAll {
		content = regex.ReplaceAllStringFunc(content, func(str string) string {
			return valueFunc(str)
		})
	} else {
		parts := regex.FindStringSubmatch(content)
		if len(parts) > 0 {
			content = strings.Replace(content, parts[0], valueFunc(parts...), 1)
		}
	}
	Write(path, content)
}

func ModContent(path string, regStr string, isAll bool, value string) {
	ModifyContent(path, regStr, isAll, func(matchStr ...string) string {
		return value
	})
}

func ModifyContent(path string, regStr string, isAll bool, valueFunc func(matchStr ...string) string) {
	content := Read(path)
	contentBreak := "\n"
	if strings.Contains(content, "\r\n") {
		contentBreak = "\r\n"
	}
	regex := regexp.MustCompile(regStr)
	lines := strings.Split(content, contentBreak)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		match := regex.MatchString(line)
		if !match {
			continue
		}
		parts := regex.FindStringSubmatch(line)
		if len(parts) == 0 {
			continue
		}
		lines[i] = strings.Replace(line, parts[1], valueFunc(parts...), 1)
		if !isAll {
			break
		}
	}
	Write(path, arrayToStr(lines, contentBreak))
}

func arrayToStr(lstStr []string, separator string) string {
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
