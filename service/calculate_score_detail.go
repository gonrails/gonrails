package service

import (
	"kalista/models"
	"kalista/utils/common"
	"time"
)

func CalculateScoreDetail(bet_order_id uint, event_type string) {

	var betOrder *models.BetOrder
	t := time.Now()

	betOrder.FindById(bet_order_id)

	models.NewScoreDetail(
		t.Format(common.DateFormat),
		models.Game{},
		models.Tenant{},
	)
}
