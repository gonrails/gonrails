/*
Author: 曾涛
Desc:   这里用来记录 Score 的所有流动明细、成单 - 退单（就是两笔记录）
Date:   2019-05-07
Email:  zengtao@risewinter.com
*/

package models

import (
	"github.com/jinzhu/gorm"
)

/**
* * belongs_to Game
* * belongs_to Tenant
* * belongs_to Scoreable (polymorphic) Gorm 不支持
* TODO newScoreDetail() 不一定会成功 所以需要加入错误处理逻辑
* ! 每一个 Detail 记录的创建都会对应的去创建或者更新相应的统计记录 （所以直接用 AfterCreate）理论上创建之后不能被修改、所以任何更新操作要失败
 */

// ScoreDetail 久远来看应该是一个多态关联, ScoreableType ScoreableID
// ScoreDetail 会更新当日的统计数据 - StatisticReports
// Order -> ScoreDetail -> StatisticReports
type ScoreDetail struct {
	ID            uint   `gorm:"primary_key"`
	ScoreableType string `gorm:"not null;unique_index:idx_for_score_details;"`
	ScoreableID   uint   `gorm:"not null;unique_index:idx_for_score_details;"`
	Date          string `gorm:"type:varchar(10);"`
	Game          Game
	GameID        uint `gorm:"index:idx_game_details"`
	Tenant        Tenant
	TenantID      uint    `gorm:"index:idx_tenant_details"`
	Reason        string  `gorm:"column:reason;unique_index:idx_for_score_details;"`
	Amount        float32 `sql:"type:decimal(15, 5);"`
}

// StatisticReport 这里会查找对应的统计数据（而且，统计数据只能在这里被更新）
func (sd *ScoreDetail) StatisticReport() *StatisticReports {
	return nil
}

func (sd *ScoreDetail) AfterCreate(scope *gorm.Scope) (err error) {
	sd.updateReport()
	return nil
}

// TableName - Customize db's tablename
func (ScoreDetail) TableName() string {
	return "score_details"
}

func newScoreDetail(scoreableType, date, reason string, gameID, tenantID, scoreableID uint, amount float32) error {
	if err := db.Create(&ScoreDetail{
		ScoreableType: scoreableType,
		ScoreableID:   scoreableID,
		Date:          date,
		GameID:        gameID,
		TenantID:      tenantID,
		Reason:        reason,
		Amount:        amount,
	}).Error; err != nil {
		return err
	}
	return nil
}

/**
 * ! 如果 Amount > 0, 那么统计收入情况、如果 Amount < 0 那么统计支出情况
 */
func (sd *ScoreDetail) updateReport() {
	var report = sd.findReport()
	if sd.Amount > 0 {
		report.Turnover = report.Turnover + uint(sd.Amount)
	} else {
		report.Payback = report.Payback - sd.Amount
	}
	db.Save(report)
}

// StatisticReports UniqueIndex (Date GameID TenantID TopicType)
func (sd *ScoreDetail) findReport() *StatisticReports {
	report := &StatisticReports{Date: sd.Date, GameID: sd.GameID, TenantID: sd.TenantID}
	db.FirstOrCreate(report, *report)
	return report
}
