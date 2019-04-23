package tenant

type baseTenantSerializer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	ApiKey string `json:"api_key"`
}
