package Log

import (
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
	return DataUtil.JsonArrayToObjects(jsonArrayStr, &Log{}).([]*Log)
}

func JsonToObject(jsonStr string) *Log {
	return DataUtil.JsonToObject(jsonStr, &Log{}).(*Log)
}

func ArrayToObjects(arrayData []map[string]any) []*Log {
	return DataUtil.ArrayToObjects(arrayData, &Log{}).([]*Log)
}

func MapToObject(mapData map[string]any) *Log {
	return DataUtil.MapToObject(mapData, &Log{}).(*Log)
}

func ObjectToMap(data *Log) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func ObjectsToArray(lstData []*Log) []map[string]any {
	return DataUtil.ObjectsToArray(lstData)
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
