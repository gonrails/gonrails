package tenant

type AdminTenantIndexSerializer struct {
	baseTenantSerializer
}

func (s *AdminTenantIndexSerializer) Serialize(v interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id": "abc",
	}
}
