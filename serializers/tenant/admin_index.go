package tenant

type AdminTenantIndexSerializer struct {
	*BaseTenantSerializer
}

func (s *AdminTenantIndexSerializer) Serialize(v interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id": "abc",
	}
}
