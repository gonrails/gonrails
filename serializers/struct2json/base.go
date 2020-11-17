package struct2json

import (
	"errors"
	"reflect"
	"time"

	"github.com/gonrails/gonrails/serializers"
	"github.com/sirupsen/logrus"
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

func getMapValue(v reflect.Value, t reflect.Type) (i interface{}) {

	if reflect.Ptr == v.Kind() {
		v = v.Elem()
	}

	if reflect.Invalid == v.Kind() {
		return nil
	}

	modelType := reflect.TypeOf((*serializers.Serializer)(nil)).Elem()

	if t.Implements(modelType) {
		method := v.MethodByName("Serialize")
		return method.Call([]reflect.Value{v})
	}

	if reflect.Struct == v.Kind() {
		ans, _ := Convert(v.Interface())
		return ans
	}

	if reflect.Slice == v.Kind() {
		var ans []interface{}
		for i := 0; i < v.Len(); i++ {
			t := v.Index(i).Type()
			logrus.Println(t.Name())
			value := getMapValue(v.Index(i), v.Index(i).Type())
			ans = append(ans, value)
		}
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
			subMap := getMapValue(value, filed.Type).(map[string]interface{})
			for k, v := range subMap {
				ansMap[k] = v
			}
		} else {
			switch filed.Type.String() {
			case "time.Time":
				ansMap[getKey(&filed)] = value.Interface().(time.Time)
			default:
				ansMap[getKey(&filed)] = getMapValue(value, filed.Type)
			}

		}
	}

	return ansMap, nil
}
