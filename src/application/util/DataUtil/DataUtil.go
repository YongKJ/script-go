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

func ArrayToObjects(arrayData []map[string]any, class any) any {
	classes := make([]any, len(arrayData))
	for i := 0; i < len(arrayData); i++ {
		obj := deepCopy(class)
		classes[i] = getObject(arrayData[i], &obj)
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

func deepCopy(src any) any {
	if src == nil {
		return nil
	}

	srcValue := reflect.ValueOf(src)
	//srcType := reflect.TypeOf(src)

	// 如果是指针，则需要解引用
	for srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}

	// 根据源值创建一个新的目标值
	cpy := reflect.New(srcValue.Type()).Elem()

	fmt.Println(srcValue.Kind())
	switch srcValue.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Struct:
		fmt.Println(srcValue.NumField())
		for i := 0; i < srcValue.NumField(); i++ {
			srcFiledValue := srcValue.Field(i)
			srcFiledValue = reflect.NewAt(srcFiledValue.Type(), unsafe.Pointer(srcFiledValue.UnsafeAddr())).Elem()
			fmt.Println(srcFiledValue)
			fmt.Println(srcFiledValue.Type())
			fmt.Println(srcFiledValue.Type().Name())

			cpyFieldValue := cpy.Field(i)
			cpyFieldValue = reflect.NewAt(cpyFieldValue.Type(), unsafe.Pointer(cpyFieldValue.UnsafeAddr())).Elem()
			fmt.Println(cpyFieldValue)
			fmt.Println(cpyFieldValue.Type())
			fmt.Println(cpyFieldValue.Type().Name())

			// 递归复制每一个字段
			if srcValue.Field(i).Kind() == reflect.Struct {
				cpyFieldValue.Set(reflect.ValueOf(deepCopy(srcFiledValue)))
			} else {
				fmt.Println(reflect.ValueOf(srcFiledValue))
				cpyFieldValue.Set(reflect.ValueOf(srcFiledValue.Interface()))
			}
		}
	default:
		cpy.Set(srcValue)
	}

	return cpy.Interface()
}
