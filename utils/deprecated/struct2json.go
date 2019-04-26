/*
Author: 曾涛
Desc:   Struct 转化成为 map[string]interface{}
Date:   2019-04-24
Email:  zengtao@risewinter.com
*/

/*
ChangeLogs:

---------------------------------------------

Date:   2019-04-25
Author: 曾涛
Desc:   增加嵌套和空指针的支持 (递归调用) 支持 noroot tag

---------------------------------------------
*/

package deprecated

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

func getNoroot(f *reflect.StructField) bool {
	if _, ok := f.Tag.Lookup("noroot"); ok {
		return true
	}
	return false
}

func getMapValue(v reflect.Value) (i interface{}) {

	if reflect.Ptr == v.Kind() {
		v = v.Elem()
	}

	if reflect.Invalid == v.Kind() {
		return nil
	}

	if reflect.Struct == v.Kind() {
		ans, _ := Convert(v.Interface())
		return ans
	}
	return v.Interface()
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

		if getNoroot(&filed) {
			subMap := getMapValue(value).(map[string]interface{})
			for k, v := range subMap {
				ansMap[k] = v
			}
		} else {
			ansMap[getKey(&filed)] = getMapValue(value)
		}
	}

	return ansMap, nil
}
