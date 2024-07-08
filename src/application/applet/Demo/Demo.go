package Demo

import (
	"encoding/base64"
	lzstring "github.com/daku10/go-lz-string"
	"script-go/src/application/deploy/pojo/dto/DemoTest"
	"script-go/src/application/deploy/pojo/po/Script"
	"script-go/src/application/pojo/dto/Log"
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
}

func Run() {
	demo := newDemo()
	demo.test3()
	//demo.test2()
	//demo.test1()
	//demo.test()
}
