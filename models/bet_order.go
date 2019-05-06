package models

import "time"

// belongs_to Tenant
// has_many ScoreDetails

type BetOrder struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Tenant   Tenant
	TenantID uint

	Amount int     `gorm:"column:amount;not null;default: 0"`
	Odd    float32 `sql:"type:decimal(15, 5)"`

	Status        int
	TenantUserID  uint          `gorm:"column:tenant_customer_no"`
	TenantOrderNo string        `gorm:"column:tenant_order_no"`
	ScoreDetails  []ScoreDetail `gorm:"polymorphic:Scoreable;"` // 一个订单可能关联多个 ScoreDetail
}

func (bet_order *BetOrder) FindById(id uint) {
	db.Where(&BetOrder{
		ID: id,
	}).First(bet_order)
}

func (BetOrder) TableName() string {
	return "bet_orders"
}
