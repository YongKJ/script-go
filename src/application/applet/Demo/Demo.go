package Demo

import (
	"encoding/base64"
	lzstring "github.com/daku10/go-lz-string"
	"log"
	"os"
	"path/filepath"
	"script-go/src/application/deploy/pojo/dto/DemoTest"
	"script-go/src/application/deploy/pojo/dto/TestDemo"
	"script-go/src/application/deploy/pojo/po/Script"
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/GenUtil"
	"script-go/src/application/util/LogUtil"
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

	lstObjData := DemoTest.ArrayToObjects(lstData)
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
		log.Println(err)
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "appPath", appPath))

	execPath, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "execPath", execPath))
}

func (d *Demo) test7() {
	execPath, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "execPath", execPath))

	yamlPath := GenUtil.GetYamlByContent(execPath)
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test6", "yamlPath", yamlPath))
}

func (d *Demo) test8() {
	msg := GenUtil.GetValue("msg")
	LogUtil.LoggerLine(Log.Of("ApplicationTest", "test8", "msg", msg))
}

func Run() {
	demo := newDemo()
	demo.test8()
	//demo.test7()
	//demo.test6()
	//demo.test5()
	//demo.test4()
	//demo.test3()
	//demo.test2()
	//demo.test1()
	//demo.test()
}
