package report

import (
	"kalista/models"

	"github.com/w-zengtao/struct2json"
)

type AdminRevenueReportIndexSerializer struct {
	ID       uint    `json:"id"`
	Date     string  `json:"date"`
	Payback  float32 `json:"payback"`
	Turnover uint    `json:"turnover"`
}

func (s *AdminRevenueReportIndexSerializer) Serialize(v interface{}) map[string]interface{} {

	if obj, ok := v.(*models.RevenueReport); ok {
		s = &AdminRevenueReportIndexSerializer{
			ID:   obj.ID,
			Date: obj.Date,
		}

		ans, _ := struct2json.Convert(s)
		return ans
	}

	return map[string]interface{}{"error": "params passed to method Serialize(v interface{}) is not a Report"}
}
