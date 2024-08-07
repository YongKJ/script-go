package DataUtil

import (
	"encoding/json"
	"log"
	"reflect"
	"unsafe"
)

func JsonArrayToMaps(jsonStr string) []map[string]any {
	var arrayData []map[string]any
	return JsonArrayToObjects(jsonStr, arrayData).([]map[string]any)
}

func JsonArrayToObjects(jsonStr string, class any) any {
	err := json.Unmarshal([]byte(jsonStr), &class)
	if err != nil {
		log.Println("[DateUtil] JsonArrayToObjects -> json.Unmarshal: ", err)
	}
	return class
}

func JsonToMap(jsonStr string) map[string]any {
	var mapData map[string]any
	err := json.Unmarshal([]byte(jsonStr), &mapData)
	if err != nil {
		log.Println("[DateUtil] JsonToMap -> json.Unmarshal: ", err)
	}
	return mapData
}

func JsonToObject(jsonStr string, class any) any {
	err := json.Unmarshal([]byte(jsonStr), class)
	if err != nil {
		log.Println("[DateUtil] JsonToObject -> json.Unmarshal: ", err)
	}
	return class
}

func ObjectsToJsonArray[A ~[]E, E any](classes A) string {
	bytes, err := json.Marshal(classes)
	if err != nil {
		log.Println("[DateUtil] ObjectsToJsonArray -> json.Marshal: ", err)
	}
	return string(bytes)
}

func ObjectsToMaps[A ~[]E, E any](classes A) []map[string]any {
	jsonStr := ObjectsToJsonArray(classes)
	return JsonArrayToMaps(jsonStr)
}

func ObjectToJson(class any) string {
	bytes, err := json.Marshal(class)
	if err != nil {
		log.Println("[DateUtil] ObjectToString -> json.Marshal: ", err)
	}
	return string(bytes)
}

func ObjectToMap(class any) map[string]any {
	jsonStr := ObjectToJson(class)
	return JsonToMap(jsonStr)
}

func MapsToJsonArray(lstData []map[string]any) string {
	bytes, err := json.Marshal(lstData)
	if err != nil {
		log.Println("[DateUtil] MapsToJsonArray -> json.Marshal: ", err)
	}
	return string(bytes)
}

func MapsToObjects(lstData []map[string]any, class any) any {
	jsonStr := MapsToJsonArray(lstData)
	return JsonArrayToObjects(jsonStr, class)
}

func MapToJson(mapData map[string]any) string {
	bytes, err := json.Marshal(mapData)
	if err != nil {
		log.Println("[DateUtil] MapToJson -> json.Marshal: ", err)
	}
	return string(bytes)
}

func MapToObject(mapData map[string]any, class any) any {
	jsonStr := MapToJson(mapData)
	return JsonToObject(jsonStr, class)
}

func ObjectsToArray[A ~[]E, E any](classes A) []map[string]any {
	jsonStr := ObjectsToJsonArray(classes)
	return JsonArrayToMaps(jsonStr)
}

func ObjectsToArrayOld[A ~[]E, E any](classes A) []map[string]any {
	length := len(classes)
	lstData := make([]map[string]any, length)
	for i := 0; i < length; i++ {
		lstData[i] = getMap(classes[i])
	}
	return lstData
}

func MapToObjectOld(mapData map[string]any, class any) any {
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
			refValue = reflect.ValueOf(MapToObjectOld(fieldValue.(map[string]any), cpyValues))
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
