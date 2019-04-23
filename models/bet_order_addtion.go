package models

// belongs_to Game
// beongs_to BetOrder

type BetOrderAddition struct {
	ID         uint `gorm:"primary_key"`
	Game       Game
	GameID     uint
	BetOrder   BetOrder
	BetOrderID uint
}

func (BetOrderAddition) TableName() string {
	return "bet_order_additions"
}
