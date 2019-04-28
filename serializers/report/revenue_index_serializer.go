package report

type AdminRevenueReportIndexSerializer struct {
	ID   uint   `json:"id"`
	Date string `json:"date"`
}

func (s *AdminRevenueReportIndexSerializer) Serialize(v interface{}) map[string]interface{} {

}
