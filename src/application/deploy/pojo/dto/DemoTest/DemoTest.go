package DemoTest

import (
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/GenUtil"
)

type DemoTest struct {
	id  int
	msg string
}

func newDemoTest(id int, msg string) *DemoTest {
	return &DemoTest{id: id, msg: msg}
}

func Of(id int, msg string) *DemoTest {
	return newDemoTest(id, msg)
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

func ObjectToMap(script *DemoTest) map[string]any {
	return DataUtil.ObjectToMap(script)
}

func ObjectsToArray(scripts []*DemoTest) []map[string]any {
	return DataUtil.ObjectsToArray(GenUtil.ArraysToAny(scripts))
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
