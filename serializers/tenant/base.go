package tenant

// BaseTenantSerializer is the basic tenant serializer info
type BaseTenantSerializer struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
}
