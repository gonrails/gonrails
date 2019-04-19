package models

type BetOrderAddition struct {
	Game       Game
	GameID     uint
	BetOrder   BetOrder
	BetOrderID uint
}

func (BetOrderAddition) TableName() string {
	return "bet_order_additions"
}
