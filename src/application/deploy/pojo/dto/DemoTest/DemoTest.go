package DemoTest

import (
	"script-go/src/application/deploy/pojo/dto/TestDemo"
	"script-go/src/application/util/DataUtil"
)

type DemoTest struct {
	id       int
	msg      string
	testDemo *TestDemo.TestDemo
}

func newDemoTest(id int, msg string, testDemo *TestDemo.TestDemo) *DemoTest {
	return &DemoTest{id: id, msg: msg, testDemo: testDemo}
}

func Of(id int, msg string, testDemo *TestDemo.TestDemo) *DemoTest {
	return newDemoTest(id, msg, testDemo)
}

func JsonArrayToObjects(jsonArrayStr string) []*DemoTest {
	return DataUtil.JsonArrayToObjects(jsonArrayStr, &DemoTest{}).([]*DemoTest)
}

func JsonToObject(jsonStr string) *DemoTest {
	return DataUtil.JsonToObject(jsonStr, &DemoTest{}).(*DemoTest)
}

func ArrayToObjects(arrayData []map[string]any) []*DemoTest {
	return DataUtil.ArrayToObjects(arrayData, &DemoTest{}).([]*DemoTest)
}

func MapToObject(mapData map[string]any) *DemoTest {
	return DataUtil.MapToObject(mapData, &DemoTest{}).(*DemoTest)
}

func ObjectToMap(data *DemoTest) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func ObjectsToArray(lstData []*DemoTest) []map[string]any {
	return DataUtil.ObjectsToArray(lstData)
}

func (d *DemoTest) Id() int {
	return d.id
}

func (d *DemoTest) SetId(id int) {
	d.id = id
}

func (d *DemoTest) Msg() string {
	return d.msg
}

func (d *DemoTest) SetMsg(msg string) {
	d.msg = msg
}

func (d *DemoTest) TestDemo() *TestDemo.TestDemo {
	return d.testDemo
}

func (d *DemoTest) SetTestDemo(testDemo *TestDemo.TestDemo) {
	d.testDemo = testDemo
}
