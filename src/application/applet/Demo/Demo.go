package Demo

import (
	"bufio"
	"encoding/base64"
	"fmt"
	lzstring "github.com/daku10/go-lz-string"
	"github.com/jinzhu/copier"
	"os"
	"path/filepath"
	"reflect"
	"script-go/src/application/deploy/pojo/dto/DemoTest"
	"script-go/src/application/deploy/pojo/dto/TestDemo"
	"script-go/src/application/deploy/pojo/po/Script"
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/FileUtil"
	"script-go/src/application/util/GenUtil"
	"script-go/src/application/util/LogUtil"
	"strings"
)

type Demo struct {
}

func newDemo() *Demo {
	return &Demo{}
}

func (d *Demo) test() {
	LogUtil.LoggerLine(Log.Of("Demo", "test", "msg", "Hello world!"))
}

func (d *Demo) test1() {
	content := "Hello world"
	contentBase64 := base64.StdEncoding.EncodeToString([]byte(content))
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "contentBase64", contentBase64))
	contentBase64 = base64.StdEncoding.EncodeToString([]byte(contentBase64))
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "contentBase64", contentBase64))

	compressedStr, err := lzstring.CompressToEncodedURIComponent(contentBase64)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "err", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "compressedStr", compressedStr))

	decompressedStr, err := lzstring.DecompressFromEncodedURIComponent(compressedStr)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "err", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "decompressedStr", decompressedStr))

	message, err := base64.StdEncoding.DecodeString(decompressedStr)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "err", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "string(message)", string(message)))
	message, err = base64.StdEncoding.DecodeString(string(message))
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "err", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test1", "string(message)", string(message)))
}

func (d *Demo) test2() {
	lstScript := Script.Gets()
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "Test2", "lstScript", lstScript))
}

func (d *Demo) test3() {
	lstData := make([]map[string]any, 2)
	lstData[0] = map[string]any{
		"id":  1,
		"msg": "Hello world!",
	}
	lstData[1] = map[string]any{
		"id":  2,
		"msg": "Demo test.",
	}

	lstObjData := DemoTest.MapsToObjects(lstData)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test3", "lstObjData", lstObjData))

	objData := DataUtil.DeepCopy(lstObjData[0])
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test3", "objData", objData))
}

func (d *Demo) test4() {
	mapData := map[string]any{
		"id":  1,
		"msg": "Hello world!",
	}

	testDemo := TestDemo.MapToObject(mapData)
	demoTest := DemoTest.Of(2, "Demo test.", testDemo)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test4", "demoTest", demoTest))

	cpyDemoTest := DataUtil.DeepCopy(demoTest).(*DemoTest.DemoTest)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test4", "cpyDemoTest", cpyDemoTest))

}

func (d *Demo) test5() {
	lstData := make([]map[string]any, 2)
	lstData[0] = map[string]any{
		"id":  1,
		"msg": "Hello world!",
	}
	lstData[1] = map[string]any{
		"id":  2,
		"msg": "Demo test.",
	}

	//lstObjData := TestDemo.ArrayToObjects(lstData)
	//LogUtil.LoggerLine(Log.Of("ApplicationTest", "test3", "lstObjData", lstObjData))

	//objData := DataUtil.DeepCopy(lstObjData[0])
	//LogUtil.LoggerLine(Log.Of("ApplicationTest", "test3", "objData", objData))

	demoTest := DemoTest.MapToObject(lstData[0])
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test3", "demoTest", demoTest))
}

func (d *Demo) test6() {
	appPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "filepath.Abs", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "appPath", appPath))

	execPath, err := os.Executable()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "os.Executable", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "execPath", execPath))
}

func (d *Demo) test7() {
	execPath, err := os.Executable()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "test7", "os.Executable", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "execPath", execPath))

	yamlPath := GenUtil.GetYamlByContent(execPath)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "yamlPath", yamlPath))
}

func (d *Demo) test8() {
	msg := GenUtil.GetValue("msg")
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test8", "msg", msg))
}

func (d *Demo) test9() {
	mapData := map[string]any{
		"id":  1,
		"msg": "Hello world!",
	}
	demoTest := DemoTest.MapToObject(mapData)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test9", "demoTest", demoTest))

	values := reflect.ValueOf(demoTest)
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	cpyDemoTest := reflect.New(values.Type()).Interface()
	err := copier.Copy(cpyDemoTest, demoTest)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApplicationTest", "test9", "copier.Copy", err))
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test9", "cpyDemoTest", cpyDemoTest))
}

func (d *Demo) test10() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取控制台输入失败:", err)
		return
	}
	text = strings.TrimSpace(text)
	fmt.Println("您输入的内容是:", text)
}

func (d *Demo) test11() {
	fmt.Print("请输入内容: ")
	params := GenUtil.ReadParams()
	fmt.Println("您输入的内容是:", params)

	fmt.Print("请输入内容: ")
	params = GenUtil.ReadParams()
	fmt.Println("您输入的内容是:", params)
}

func (d *Demo) test12() {
	appDir := FileUtil.AppDir()
	LogUtil.LoggerLine(Log.Of("Demo", "test12", "appDir", appDir))
}

func (d *Demo) test13() {
	send := make(chan string)
	go func() {
		defer close(send)

		select {
		case msg, ok := <-send:
			LogUtil.LoggerLine(Log.Of("Demo", "test13", "msg", msg))
			LogUtil.LoggerLine(Log.Of("Demo", "test13", "ok", ok))
		default:
			LogUtil.LoggerLine(Log.Of("Demo", "test13", "default", "default"))
		}
	}()
	send <- "Hello world!"
	close(send)
}

func (d *Demo) test14() {
	lstData := make([]map[string]any, 2)
	lstData[0] = map[string]any{
		"id":  1,
		"msg": "Hello world!",
		"testDemo": map[string]any{
			"id":  1,
			"msg": "Hello world!",
		},
	}
	lstData[1] = map[string]any{
		"id":  1,
		"msg": "Hello world!",
		"testDemo": map[string]any{
			"id":  1,
			"msg": "Hello world!",
		},
	}

	lstObjData := DemoTest.MapsToObjects(lstData)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test3", "lstObjData", lstObjData))
}

func (d *Demo) test15() {
	userName := "yongkj"
	password := "*Dxj1003746818"
	userNameEncode := GenUtil.GetEncode(userName)
	passwordEncode := GenUtil.GetEncode(GenUtil.GetMd5Str(password))
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test15", "userNameEncode", userNameEncode))
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test15", "passwordEncode", passwordEncode))

	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyZWZyZXNoU3RyIjoiQ3JCQ0RVRk1GVUZZRFlCS1lCeUJEQTFyQVlzQU5vZ0ZZQ3kyMkFqZ0ZvQmUyQThzUUtMUUFNK2lBNG91TnVPTUNqcG84ZUFNekErYUFJb3NBak5tUXNBdkVBIiwiZXhwIjoxNzIzMjk3NDUyLCJuYmYiOjE3MjMyMTEwNTIsImlhdCI6MTcyMzIxMTA1Mn0.cx4M73-jz4Ke6ZlGM3VKpLihESoD5YlggCpFHc5EdrI"
	refreshTokenEncode := GenUtil.GetEncode(refreshToken)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test15", "refreshTokenEncode", refreshTokenEncode))
}

func (d *Demo) test16() {
	path := "D:\\Document\\MyCodes\\Gitea\\api-go\\src\\application\\common\\module\\socketio"
	checkGoFile(path)
}

func checkGoFile(folder string) {
	files := FileUtil.List(folder)
	for _, file := range files {
		filePath := filepath.Join(folder, file)
		if FileUtil.IsFolder(filePath) {
			checkGoFile(filePath)
			continue
		}

		if !strings.HasSuffix(filePath, ".go") {
			continue
		}

		FileUtil.ModFile(filePath, "github.com/googollee/go-socket.io", true, "api-go/src/application/common/module/socketio")
	}
}

func Run() {
	demo := newDemo()
	demo.test16()
	//demo.test15()
	//demo.test14()
	//demo.test13()
	//demo.test12()
	//demo.test11()
	//demo.test10()
	//demo.test9()
	//demo.test8()
	//demo.test7()
	//demo.test6()
	//demo.test5()
	//demo.test4()
	//demo.test3()
	//demo.test2()
	//demo.test1()
	//demo.test()
}
