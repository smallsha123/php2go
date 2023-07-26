package tool

import (
	"github.com/spf13/cast"
	"net/url"
	"reflect"
)

func Display(req interface{}, v1 *url.Values) {
	reflectType := reflect.TypeOf(req).Elem()
	reflectValue := reflect.ValueOf(req).Elem()
	for i := 0; i < reflectType.NumField(); i++ {
		valueValue := reflectValue.Field(i).Interface()
		typeName := reflectType.Field(i).Name
		switch reflectValue.Field(i).Kind() {
		case reflect.Struct:
			//v := reflectValue.Field(i).Addr()
			//Display(v.Interface(), v1)
			//Display(&valueValue, v1)
		default:
			v1.Add(typeName, cast.ToString(valueValue))
		}
	}
}
