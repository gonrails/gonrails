/*
Author: 曾涛
Desc:   处理序列化的 基本接口类型 描述
Date:   2019-04-23
Email:  zengtao@risewinter.com
*/

package serializers

type Serializer interface {
	Serialize(v interface{}) map[string]interface{}
}

// Template 应该设计如此 ( template.Serialize(v))
func SingleSerializer(v interface{}, s Serializer) map[string]interface{} {
	return s.Serialize(v)
}

func CollectionSerializer(v []interface{}, template interface{}) {

}
