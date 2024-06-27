package Log

import (
	"encoding/json"
	"log"
	"script-go/src/application/util/DataUtil"
)

type Log struct {
	className  string
	methodName string
	paramName  string
	value      any
}

func newLog(className string, methodName string, paramName string, value any) *Log {
	return &Log{
		className:  className,
		methodName: methodName,
		paramName:  paramName,
		value:      value,
	}
}

func Of(className string, methodName string, paramName string, value any) *Log {
	return newLog(className, methodName, paramName, value)
}

func JsonArrayToObjects(jsonArrayStr string) []*Log {
	var arrayData []map[string]any
	err := json.Unmarshal([]byte(jsonArrayStr), &arrayData)
	if err != nil {
		log.Println(err)
	}
	return ArrayToObjects(arrayData)
}

func JsonToObject(jsonStr string) *Log {
	var mapData map[string]any
	err := json.Unmarshal([]byte(jsonStr), &mapData)
	if err != nil {
		log.Println(err)
	}
	return MapToObject(mapData)
}

func ArrayToObjects(arrayData []map[string]any) []*Log {
	length := len(arrayData)
	lstData := make([]*Log, length)
	for i := 0; i < length; i++ {
		lstData[i] = MapToObject(arrayData[i])
	}
	return lstData
}

func MapToObject(mapData map[string]any) *Log {
	return DataUtil.MapToObject(mapData, &Log{}).(*Log)
}

func ObjectToMap(log *Log) map[string]any {
	return DataUtil.ObjectToMap(log)
}

func ObjectsToArray(logs []*Log) []map[string]any {
	length := len(logs)
	lstData := make([]map[string]any, length)
	for i := 0; i < length; i++ {
		lstData[i] = ObjectToMap(logs[i])
	}
	return lstData
}

func (l *Log) Value() any {
	return l.value
}

func (l *Log) SetValue(value any) {
	l.value = value
}

func (l *Log) ParamName() string {
	return l.paramName
}

func (l *Log) SetParamName(paramName string) {
	l.paramName = paramName
}

func (l *Log) MethodName() string {
	return l.methodName
}

func (l *Log) SetMethodName(methodName string) {
	l.methodName = methodName
}

func (l *Log) ClassName() string {
	return l.className
}

func (l *Log) SetClassName(className string) {
	l.className = className
}
