package models

type baseTenantSerializer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	ApiKey string `json:"api_key"`
}

// AdminTenantIndexSerializer - GET /admin/tenants
type AdminTenantIndexSerializer struct {
	baseTenantSerializer
}

func (s *AdminTenantIndexSerializer) Serialize(v interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id": "abc",
	}
}

type AdminTenantShowSerializer struct {
	baseTenantSerializer
}

func (s *AdminTenantShowSerializer) Serialize(v interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id": "abc",
	}
}

func (t *Tenant) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id": t.ID,
	}
}
