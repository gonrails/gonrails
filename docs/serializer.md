### Serializer

* directory `serializers` helps us marshal request into desired go types.

#### package serializer

```go
package serializers

import (
	"errors"
	"reflect"
)

// Serializer is the basic interface for API Response
type Serializer interface {
	Serialize(v interface{}) map[string]interface{}
}

// SingleSerializer 处理单个对象的序列化
func SingleSerializer(s Serializer, v interface{}) map[string]interface{} {
	return s.Serialize(v)
}

// CollectionSerializer 处理集合对象的序列化
func CollectionSerializer(s Serializer, vs interface{}) ([]map[string]interface{}, error) {
	ansAry := []map[string]interface{}{}

	if reflect.Slice != reflect.TypeOf(vs).Kind() {
		return []map[string]interface{}{}, errors.New("Data must be slice type")
	}

	value := reflect.ValueOf(vs)

	for i := 0; i < value.Len(); i++ {
		ansAry = append(ansAry, s.Serialize(value.Index(i).Interface()))
	}

	return ansAry, nil
}
```
