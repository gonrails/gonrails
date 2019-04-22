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
