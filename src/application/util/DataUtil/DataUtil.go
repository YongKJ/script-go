package DataUtil

import (
	"encoding/json"
	"log"
	"reflect"
	"unsafe"
)

func JsonArrayToMaps(jsonStr string) []map[string]any {
	var arrayData []map[string]any
	err := json.Unmarshal([]byte(jsonStr), &arrayData)
	if err != nil {
		log.Println(err)
	}
	return arrayData
}

func JsonToMap(jsonStr string) map[string]any {
	var mapData map[string]any
	err := json.Unmarshal([]byte(jsonStr), &mapData)
	if err != nil {
		log.Println(err)
	}
	return mapData
}

func JsonToObject(jsonStr string, class any) any {
	var mapData map[string]any
	err := json.Unmarshal([]byte(jsonStr), &mapData)
	if err != nil {
		log.Println(err)
	}
	return MapToObject(mapData, class)
}

func ObjectToMap(class any) map[string]any {
	return getMap(class)
}

func ObjectsToArray(classes []any) []map[string]any {
	length := len(classes)
	lstData := make([]map[string]any, length)
	for i := 0; i < length; i++ {
		lstData[i] = getMap(classes[i])
	}
	return lstData
}

func MapToObject(mapData map[string]any, class any) any {
	return getObject(mapData, class)
}

func ArrayToObjects(arrayData []map[string]any, classes []any) []any {
	for i := 0; i < len(arrayData); i++ {
		getObject(arrayData[i], classes[i])
	}
	return classes
}

func getObject(mapData map[string]any, class any) any {
	values := reflect.ValueOf(class)
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	types := reflect.TypeOf(class)
	if types.Kind() == reflect.Ptr {
		types = types.Elem()
	}
	for i := 0; i < types.NumField(); i++ {
		name := types.Field(i).Name
		value := values.FieldByName(name)
		value = reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()

		typeName := types.Field(i).Type.Name()
		switch typeName {
		case "int":
			value.Set(reflect.ValueOf(int(mapData[name].(float64))))
			break
		default:
			value.Set(reflect.ValueOf(mapData[name]))
		}
	}
	return class
}

func getMap(class any) map[string]any {
	values := reflect.ValueOf(class)
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	types := reflect.TypeOf(class)
	if types.Kind() == reflect.Ptr {
		types = types.Elem()
	}
	mapData := make(map[string]any)
	for i := 0; i < types.NumField(); i++ {
		name := types.Field(i).Name
		value := values.FieldByName(name)
		value = reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()
		mapData[name] = value.Interface()
	}
	return mapData
}

func SetValue(class any, fieldKey string, fieldValue any) {
	values := reflect.ValueOf(class)
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	if values.Kind() != reflect.Struct {
		return
	}
	value := values.FieldByName(fieldKey)
	value = reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()
	value.Set(reflect.ValueOf(fieldValue))
}
