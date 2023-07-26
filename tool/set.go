package tool

import (
	"errors"
	"reflect"
)

func DiffSet(a []int64, b []int64) []int64 {
	var diff []int64
	mp := make(map[int64]int)
	for _, n := range b {
		if _, ok := mp[n]; !ok {
			mp[n] = 1
		}
	}

	for _, n := range a {
		if _, ok := mp[n]; !ok {
			diff = append(diff, n)
		}
	}

	return diff
}

// @title struct转map 返回的map键为struct的成员名
func StructToMap(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {

		data[t.Field(i).Name] = v.Field(i).Interface()

	}

	return data
}

// 将一个slice 转为map[field]value 格式
func SliceToMap(slice interface{}, fieldName string) (map[interface{}]interface{}, error) {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		return nil, errors.New("input is not a slice")
	}

	if s.Len() == 0 {
		return nil, nil
	}

	elemType := s.Index(0).Type()

	if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
		elemType = elemType.Elem()
	}

	if _, ok := elemType.FieldByName(fieldName); !ok {
		return nil, errors.New("field does not exist in struct")
	}

	result := make(map[interface{}]interface{})
	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		field := elem.FieldByName(fieldName)
		if field.Interface() == nil {
			continue
		}
		result[field.Interface()] = s.Index(i).Interface()
	}

	return result, nil
}
//切片去重 removeDuplicates
func RemoveDuplicates(arr []string) []string {
    var  result  []string
    tempMap := map[string]bool{}
    for _, item := range arr {
        if _, ok := tempMap[item]; !ok {
            tempMap[item] = true
            result = append(result, item)
        }
    }
    return result
}
