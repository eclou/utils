package utils

func MapKeys[K comparable, V interface{} | *interface{}](mp map[K]V) []K {
	var result []K
	for k := range mp {
		result = append(result, k)
	}

	return result
}

func MapValues[K comparable, V interface{} | *interface{}](mp map[K]V) []V {
	var result []V

	for _, v := range mp {
		result = append(result, v)
	}
	return result
}

func ArrayUnique[V comparable](arr []V) []V {
	var mp = make(map[V]int, 0)
	for _, v := range arr {
		mp[v] = 1
	}
	return MapKeys(mp)
}

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
