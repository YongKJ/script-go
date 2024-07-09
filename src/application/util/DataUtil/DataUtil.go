package DataUtil

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

func JsonArrayToObjects(jsonStr string, class any) any {
	var arrayData []map[string]any
	err := json.Unmarshal([]byte(jsonStr), &arrayData)
	if err != nil {
		log.Println(err)
	}
	return ArrayToObjects(arrayData, class)
}

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

func ObjectsToArray[A ~[]E, E any](classes A) []map[string]any {
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

func ArrayToObjects(arrayData []map[string]any, class any) any {
	obj := reflect.ValueOf(class)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}
	classes := getArray(class, len(arrayData))
	for i, mapData := range arrayData {
		cpyObj := DeepCopy(class)
		objData := getObject(mapData, cpyObj)
		classes.Index(i).Set(reflect.ValueOf(objData))
	}
	return classes.Interface()
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
		fieldValue, ok := mapData[name]
		if !ok {
			continue
		}
		value := values.FieldByName(name)
		value = reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()

		fmt.Println(value.Kind())
		fmt.Println(value.Type())
		fmt.Println(types.Field(i).Type)

		var refValue reflect.Value
		switch value.Kind() {
		case reflect.Int:
			if number, ok := fieldValue.(float64); ok {
				refValue = reflect.ValueOf(int(number))
			} else {
				refValue = reflect.ValueOf(fieldValue)
			}
		case reflect.Ptr:
			cpyValues := reflect.New(types.Field(i).Type).Interface()
			refValue = reflect.ValueOf(MapToObject(fieldValue.(map[string]any), cpyValues))
		default:
			refValue = reflect.ValueOf(fieldValue)
		}
		value.Set(refValue)
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

func getArray(class any, length int) reflect.Value {
	return reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(class)), length, length)
}

func DeepCopy(class any) any {
	values := reflect.ValueOf(class)
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	types := reflect.TypeOf(class)
	if types.Kind() == reflect.Ptr {
		types = types.Elem()
	}

	cpyValues := reflect.New(values.Type()).Interface()
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i).Name
		value := values.FieldByName(field)
		value = reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()

		var fieldValue any
		switch value.Kind() {
		case reflect.Struct:
			fieldValue = DeepCopy(fieldValue)
		case reflect.Ptr:
			fieldValue = DeepCopy(fieldValue)
		default:
			fieldValue = value.Interface()
		}
		SetValue(cpyValues, field, fieldValue)
	}
	return cpyValues
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
