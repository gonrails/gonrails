package models

// AdminTenantIndexSerializer - GET /admin/tenants
type AdminTenantIndexSerializer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	ApiKey string `json:"api_key"`
}

func (t *Tenant) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id": t.ID,
	}
}
