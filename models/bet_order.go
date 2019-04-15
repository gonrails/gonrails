package models

type BetOrder struct {
}

func (BetOrder) TableName() string {
	return "bet_orders"
}
