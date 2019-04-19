package models

// belongs_to Tenant
// belongs_to Game
// has_many ScoreDetails

type BetOrder struct {
	ID           uint `gorm:"primary_key"`
	Tenant       Tenant
	TenantID     uint
	Game         Game
	GameID       uint
	Amount       int           `gorm:"column:amount;not null;default: 0"`
	Odd          float32       `sql:"type:decimal(15, 5)"`
	ScoreDetails []ScoreDetail `gorm:"polymorphic:Scoreable"` // 一个订单可能关联多个 ScoreDetail
}

func (bet_order *BetOrder) FindById(id uint) {
	db.Where(&BetOrder{
		ID: id,
	}).First(bet_order)
}

func (BetOrder) TableName() string {
	return "new_bet_orders"
}
