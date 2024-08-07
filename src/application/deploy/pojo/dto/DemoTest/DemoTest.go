package DemoTest

import (
	"script-go/src/application/deploy/pojo/dto/TestDemo"
	"script-go/src/application/util/DataUtil"
)

type DemoTest struct {
	Id       int                `json:"id"`
	Msg      string             `json:"msg"`
	TestDemo *TestDemo.TestDemo `json:"testDemo"`
}

func newDemoTest(id int, msg string, testDemo *TestDemo.TestDemo) *DemoTest {
	return &DemoTest{Id: id, Msg: msg, TestDemo: testDemo}
}

func Of(id int, msg string, testDemo *TestDemo.TestDemo) *DemoTest {
	return newDemoTest(id, msg, testDemo)
}

func JsonArrayToObjects(jsonArrayStr string) []*DemoTest {
	var lstData []*DemoTest
	return *(DataUtil.JsonArrayToObjects(jsonArrayStr, &lstData).(*[]*DemoTest))
}

func JsonArrayToMaps(jsonArrayStr string) []map[string]any {
	return DataUtil.JsonArrayToMaps(jsonArrayStr)
}

func JsonToObject(jsonStr string) *DemoTest {
	return DataUtil.JsonToObject(jsonStr, &DemoTest{}).(*DemoTest)
}

func JsonToMap(jsonStr string) map[string]any {
	return DataUtil.JsonToMap(jsonStr)
}

func ObjectsToJsonArray(lstData []*DemoTest) string {
	return DataUtil.ObjectsToJsonArray(lstData)
}

func ObjectsToMaps(lstData []*DemoTest) []map[string]any {
	return DataUtil.ObjectsToMaps(lstData)
}

func ObjectToJson(data *DemoTest) string {
	return DataUtil.ObjectToJson(data)
}

func ObjectToMap(data *DemoTest) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func MapsToJsonArray(arrayData []map[string]any) string {
	return DataUtil.MapsToJsonArray(arrayData)
}

func MapsToObjects(arrayData []map[string]any) []*DemoTest {
	var lstData []*DemoTest
	return *(DataUtil.MapsToObjects(arrayData, &lstData).(*[]*DemoTest))
}

func MapToJson(mapData map[string]any) string {
	return DataUtil.MapToJson(mapData)
}

func MapToObject(mapData map[string]any) *DemoTest {
	return DataUtil.MapToObject(mapData, &DemoTest{}).(*DemoTest)
}
