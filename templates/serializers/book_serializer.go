package serializers

import (
	"github.com/w-zengtao/struct2json"
)

type BookSerializer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *BookSerializer) Serialize(v interface{}) map[string]interface{} {

	s = &BookSerializer{
		ID:   1,
		Name: "Gonrails Demo",
	}

	ans, err := struct2json.Convert(s)

	if nil == err {
		return ans
	}

	return map[string]interface{}{
		"error": "BookSerializer Serialize Error",
	}
}
