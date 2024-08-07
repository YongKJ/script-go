package Log

import (
	"script-go/src/application/util/DataUtil"
)

type Log struct {
	ClassName  string `json:"className"`
	MethodName string `json:"methodName"`
	ParamName  string `json:"paramName"`
	Value      any    `json:"value"`
}

func newLog(className string, methodName string, paramName string, value any) *Log {
	return &Log{
		ClassName:  className,
		MethodName: methodName,
		ParamName:  paramName,
		Value:      value,
	}
}

func Of(className string, methodName string, paramName string, value any) *Log {
	return newLog(className, methodName, paramName, value)
}

func JsonArrayToObjects(jsonArrayStr string) []*Log {
	var lstData []*Log
	return *(DataUtil.JsonArrayToObjects(jsonArrayStr, &lstData).(*[]*Log))
}

func JsonArrayToMaps(jsonArrayStr string) []map[string]any {
	return DataUtil.JsonArrayToMaps(jsonArrayStr)
}

func JsonToObject(jsonStr string) *Log {
	return DataUtil.JsonToObject(jsonStr, &Log{}).(*Log)
}

func JsonToMap(jsonStr string) map[string]any {
	return DataUtil.JsonToMap(jsonStr)
}

func ObjectsToJsonArray(lstData []*Log) string {
	return DataUtil.ObjectsToJsonArray(lstData)
}

func ObjectsToMaps(lstData []*Log) []map[string]any {
	return DataUtil.ObjectsToMaps(lstData)
}

func ObjectToJson(data *Log) string {
	return DataUtil.ObjectToJson(data)
}

func ObjectToMap(data *Log) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func MapsToJsonArray(arrayData []map[string]any) string {
	return DataUtil.MapsToJsonArray(arrayData)
}

func MapsToObjects(arrayData []map[string]any) []*Log {
	var lstData []*Log
	return *(DataUtil.MapsToObjects(arrayData, &lstData).(*[]*Log))
}

func MapToJson(mapData map[string]any) string {
	return DataUtil.MapToJson(mapData)
}

func MapToObject(mapData map[string]any) *Log {
	return DataUtil.MapToObject(mapData, &Log{}).(*Log)
}
