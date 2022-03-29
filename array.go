package utils

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

//receive an array of map and specified key,return the specified key's value as array
//example: [{"name":"xiaoming","sex":"male"},{"name":"xiaohua","sex":"female"}] will return ["xiaoming","xiaohua"] if the key is "name"
func ArrayColumn[K comparable, V interface{} | *interface{}](arr []map[K]V, key K) []V {
	var result []V
	for _, v := range arr {
		v2, ok := v[key]
		if ok {
			result = append(result, v2)
		}
	}

	return result
}

//receive an array of map and specified key,return the specified key's value as new map's key
//example:[{"name":"xiaoming","sex":"male"},{"name":"xiaohua","sex":"female"}] will return
//["xiaoming":{"name":"xiaoming","sex":"male"},"xiaohua":{"name":"xiaohua","sex":"female"}] if the key is "name"
func ArrayIndex[K comparable, V comparable](arr []map[K]V, key K) map[V]map[K]V {
	var result = make(map[V]map[K]V, 0)

	for _, v := range arr {
		v2, ok := v[key]
		if ok {
			result[v2] = v
		}
	}
	return result
}
