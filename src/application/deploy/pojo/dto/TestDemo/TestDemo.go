package TestDemo

import "script-go/src/application/util/DataUtil"

type TestDemo struct {
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

func newTestDemo(id int, msg string) *TestDemo {
	return &TestDemo{Id: id, Msg: msg}
}

func Of(id int, msg string) *TestDemo {
	return newTestDemo(id, msg)
}

func JsonArrayToObjects(jsonArrayStr string) []*TestDemo {
	var lstData []*TestDemo
	return DataUtil.JsonArrayToObjects(jsonArrayStr, lstData).([]*TestDemo)
}

func JsonArrayToMaps(jsonArrayStr string) []map[string]any {
	return DataUtil.JsonArrayToMaps(jsonArrayStr)
}

func JsonToObject(jsonStr string) *TestDemo {
	return DataUtil.JsonToObject(jsonStr, &TestDemo{}).(*TestDemo)
}

func JsonToMap(jsonStr string) map[string]any {
	return DataUtil.JsonToMap(jsonStr)
}

func ObjectsToJsonArray(lstData []*TestDemo) string {
	return DataUtil.ObjectsToJsonArray(lstData)
}

func ObjectsToMaps(lstData []*TestDemo) []map[string]any {
	return DataUtil.ObjectsToMaps(lstData)
}

func ObjectToJson(data *TestDemo) string {
	return DataUtil.ObjectToJson(data)
}

func ObjectToMap(data *TestDemo) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func MapsToJsonArray(arrayData []map[string]any) string {
	return DataUtil.MapsToJsonArray(arrayData)
}

func MapsToObjects(arrayData []map[string]any) []*TestDemo {
	return DataUtil.MapsToObjects(arrayData, &TestDemo{}).([]*TestDemo)
}

func MapToJson(mapData map[string]any) string {
	return DataUtil.MapToJson(mapData)
}

func MapToObject(mapData map[string]any) *TestDemo {
	return DataUtil.MapToObject(mapData, &TestDemo{}).(*TestDemo)
}
