package models

type BetOrder struct {
	ID       uint `gorm:"primary_key"`
	Tenant   Tenant
	TenantID uint
	Amount   int     `gorm:"column:amount;not null;default: 0"`
	Odd      float32 `sql:"type:decimal(15, 5)"`
}

func (bet_order *BetOrder) FindById(id uint) {
	db.Where(&BetOrder{
		ID: id,
	}).First(bet_order)
}

func (BetOrder) TableName() string {
	return "new_bet_orders"
}

func (BetOrder) Game() {

}
