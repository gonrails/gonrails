package report

import (
	"kalista/models"

	"github.com/w-zengtao/struct2json"
)

type AdminStatisticReportIndexSerializer struct {
	ID                uint    `json:"id"`
	Date              string  `json:"date"`
	Payback           float32 `json:"payback"`
	Turnover          uint    `json:"turnover"`
	SeriesCount       uint    `json:"series_count"`
	TopicCount        uint    `json:"topic_count"`
	BetUserCount      uint    `json:"bet_user_count"`
	FreshBetUserCount uint    `json:"fresh_bet_user_count"`
	BetOrderCount     uint    `json:"bet_order_count"`
}

func (s *AdminStatisticReportIndexSerializer) Serialize(v interface{}) map[string]interface{} {

	if obj, ok := v.(*models.StatisticReports); ok {
		s = &AdminStatisticReportIndexSerializer{
			ID:                obj.ID,
			Date:              obj.Date,
			Payback:           obj.Payback,
			Turnover:          obj.Turnover,
			SeriesCount:       obj.SeriesCount,
			BetUserCount:      obj.BetUserCount,
			FreshBetUserCount: obj.FreshBetUserCount,
			BetOrderCount:     obj.BetOrderCount,
		}

		ans, _ := struct2json.Convert(s)
		return ans
	}

	return map[string]interface{}{"error": "params passed to method Serialize(v interface{}) is not a Report"}
}
