package TestDemo

import "script-go/src/application/util/DataUtil"

type TestDemo struct {
	id  int
	msg string
}

func newTestDemo(id int, msg string) *TestDemo {
	return &TestDemo{id: id, msg: msg}
}

func Of(id int, msg string) *TestDemo {
	return newTestDemo(id, msg)
}

func JsonArrayToObjects(jsonArrayStr string) []*TestDemo {
	return DataUtil.JsonArrayToObjects(jsonArrayStr, &TestDemo{}).([]*TestDemo)
}

func JsonToObject(jsonStr string) *TestDemo {
	return DataUtil.JsonToObject(jsonStr, &TestDemo{}).(*TestDemo)
}

func ArrayToObjects(arrayData []map[string]any) []*TestDemo {
	return DataUtil.ArrayToObjects(arrayData, &TestDemo{}).([]*TestDemo)
}

func MapToObject(mapData map[string]any) *TestDemo {
	return DataUtil.MapToObject(mapData, &TestDemo{}).(*TestDemo)
}

func ObjectToMap(data *TestDemo) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func ObjectsToArray(lstData []*TestDemo) []map[string]any {
	return DataUtil.ObjectsToArray(lstData)
}

func (t *TestDemo) Id() int {
	return t.id
}

func (t *TestDemo) SetId(id int) {
	t.id = id
}

func (t *TestDemo) Msg() string {
	return t.msg
}

func (t *TestDemo) SetMsg(msg string) {
	t.msg = msg
}
