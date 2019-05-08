/*
Author: 曾涛
Desc:   BetOrder的逻辑（包括上下文） 姑苏城外寒山寺，夜半钟声到客船
Date:   2019-05-06
Email:  zengtao@risewinter.com
*/

package models

import (
	"kalista/utils/common"
	"time"
)

// BetOrder - 这里用来收集用户对某个 Topic 的 某些 Option 发生概率的判断，作为预测模型一个重要维度
/**
 * ! Order 需要有创建时间、结算时间（这个时间也可以是清盘时间、由状态决定）
 * ! 这边是通过消息通信来触发我们的计算、也就是订单的创建、结算、清算、回滚、取消等操作都发生在其他的系统 & 发生之后异步到这边做统计行为
 * * belongs_to Tenant
 * * has_many ScoreDetails
 */
type BetOrder struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time // 这里就是订单的创建时间
	UpdatedAt time.Time
	CheckedAt time.Time // 这里就是订单的结算时间、如果状态是清盘、那么这里就是清盘时间

	Tenant   Tenant
	TenantID uint

	Amount int     `gorm:"column:amount;not null;default: 0"`
	Odd    float32 `sql:"type:decimal(15, 5)"`

	Status        int
	TenantUserID  uint          `gorm:"column:tenant_customer_no"`
	TenantOrderNo string        `gorm:"column:tenant_order_no"`
	ScoreDetails  []ScoreDetail `gorm:"polymorphic:Scoreable;"` // 一个订单可能关联多个 ScoreDetail
}

// payback 计算本虚拟 Order 需要偿付的成本
func (bo *BetOrder) payback() float32 {
	return 0.0
}

func (bo *BetOrder) buildScoreDetail(date, reason string, amount float32) {
	newScoreDetail(bo.TableName(), date, reason, 1, bo.TenantID, bo.ID, amount)
}

// HandleCreate - 处理 Created 的消息
func (bo *BetOrder) HandleCreate() {
	bo.buildScoreDetail(bo.CreatedAt.Format(common.DateFormat), "Created", float32(bo.Amount))
}

// HandleCheck - 处理 Check 的消息
func (bo *BetOrder) HandleCheck() {
	bo.buildScoreDetail(bo.CheckedAt.Format(common.DateFormat), "Checked", bo.payback())
}

// HandleCancel - 处理 Cancel 的消息
func (bo *BetOrder) HandleCancel() {
	bo.buildScoreDetail(bo.CheckedAt.Format(common.DateFormat), "Canceld", -float32(bo.Amount))
}

func (bo *BetOrder) HandleRollback() {
}

func (bet_order *BetOrder) FindById(id uint) {
	db.Where(&BetOrder{
		ID: id,
	}).First(bet_order)
}

func (BetOrder) TableName() string {
	return "bet_orders"
}
