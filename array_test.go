package utils

import "testing"

func TestMapKeys(t *testing.T) {
	var mp = map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	keys := MapKeys(mp)
	t.Log(keys)
	var mp2 = map[string]interface{}{
		"one":   1,
		"two":   "two",
		"three": []int{1, 12, 23},
	}
	keys2 := MapKeys(mp2)
	t.Log(keys2)
}

func TestMapValues(t *testing.T) {
	var mp = map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	keys := MapValues(mp)
	t.Log(keys)
	var mp2 = map[string]interface{}{
		"one":   1,
		"two":   "two",
		"three": []int{1, 12, 23},
	}
	keys2 := MapValues(mp2)
	t.Log(keys2)
}

func TestArrayUnique(t *testing.T) {
	var arr = []int{1, 2, 3, 2, 10, 15, 20, 15}

	uArr := ArrayUnique(arr)
	if len(uArr) != 6 {
		t.FailNow()
	} else {
		t.Log(uArr)
	}
}

type Person struct {
	Name string
	Sex  string
}

func TestArrayColumn(t *testing.T) {
	var arr = []*map[string]string{
		&map[string]string{"name": "xiaoming", "sex": "male"},
		&map[string]string{"name": "xiaodong", "sex": "female"},
	}

	nArr := ArrayColumn(arr, "name")
	t.Log(nArr)

	var arr2 = []Person{
		Person{Name: "xiaoming", Sex: "male"},
		Person{Name: "xiaodong", Sex: "femal"},
	}

	sArr := ArrayColumn(arr2, "Name")
	t.Log(sArr)
}

func TestArrayIndex(t *testing.T) {
	var arr = []map[string]string{
		map[string]string{"name": "xiaoming", "sex": "male"},
		map[string]string{"name": "xiaodong", "sex": "female"},
	}

	nArr := ArrayIndex(arr, "name")
	t.Log(nArr)

	var arr2 = []*Person{
		&Person{Name: "xiaoming", Sex: "male"},
		&Person{Name: "xiaodong", Sex: "femal"},
	}

	sArr := ArrayIndex(arr2, "Name")
	t.Log(sArr)
}
