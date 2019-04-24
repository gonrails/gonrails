/*
Author: 曾涛
Desc:   Struct 转化成为 map[string]interface{}
Date:   2019-04-24
Email:  zengtao@risewinter.com
*/

package struct2json

import (
	"errors"
	"reflect"
)

func getType(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	if reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	return t
}

func getValue(v interface{}) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(v))
}

func getKey(f *reflect.StructField) string {
	var key = f.Name
	if json, ok := f.Tag.Lookup("json"); ok {
		key = json
	}
	return key
}

// Convert converts a struct to a common JSON
// Convert should be recursive
func Convert(struc interface{}) (map[string]interface{}, error) {
	ansMap := make(map[string]interface{})

	sType := getType(struc)
	sVal := getValue(struc)

	if reflect.Struct != sType.Kind() {
		return nil, errors.New("v is not a struct or a pointer to a struct")
	}

	for i := 0; i < sType.NumField(); i++ {

		filed := sType.Field(i)
		value := sVal.FieldByName(filed.Name)
		ansMap[getKey(&filed)] = value.Interface()
	}

	return ansMap, nil
}
