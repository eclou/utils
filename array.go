package utils

import (
	"reflect"
)

//return map keys as array
func MapKeys[K comparable, V interface{} | *interface{}](mp map[K]V) []K {
	var result []K
	for k := range mp {
		result = append(result, k)
	}

	return result
}

//return map values as array
func MapValues[K comparable, V interface{} | *interface{}](mp map[K]V) []V {
	var result []V

	for _, v := range mp {
		result = append(result, v)
	}
	return result
}

//remove array repeated value
func ArrayUnique[V comparable](arr []V) []V {
	var mp = make(map[V]int, 0)
	for _, v := range arr {
		mp[v] = 1
	}
	return MapKeys(mp)
}

func columnValue[K comparable](v reflect.Value, key K) any {
	var result any
	switch v.Kind() {
	case reflect.Map:
		result = v.MapIndex(reflect.ValueOf(key)).Interface()
	case reflect.Struct:
		k2, ok := (reflect.ValueOf(key).Interface()).(string)
		if ok {
			result = v.FieldByName(k2).Interface()
		}
	case reflect.Ptr:
		result = columnValue(reflect.ValueOf(v.Elem().Interface()), key)
	}
	return result
}

//return the specified key's value as array
//example: [{"name":"xiaoming","sex":"male"},{"name":"xiaohua","sex":"female"}] will return ["xiaoming","xiaohua"] if the key is "name"
func ArrayColumn[K comparable, V any](arr []V, key K) []any {
	var result []any
	for _, v := range arr {
		v2 := columnValue(reflect.ValueOf(v), key)
		if v2 != nil {
			result = append(result, v2)
		}
	}
	return result
}

//return the specified key's value as new map's key
//example:[{"name":"xiaoming","sex":"male"},{"name":"xiaohua","sex":"female"}] will return
//["xiaoming":{"name":"xiaoming","sex":"male"},"xiaohua":{"name":"xiaohua","sex":"female"}] if the key is "name"
func ArrayIndex[K comparable, V any](arr []V, key K) map[any]V {
	var result = make(map[any]V, 0)

	for _, v := range arr {
		v2 := columnValue(reflect.ValueOf(v), key)
		if v2 != nil {
			result[v2] = v
		}
	}
	return result
}
