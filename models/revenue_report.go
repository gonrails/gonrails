/*
Author: 曾涛
Desc:   模型描述、平台的收益报告 南朝四百八十寺，多少楼台烟雨中
Date:   2019-04-26
Email:  zengtao@risewinter.com
*/

/*
商户
游戏
日期
营收（如果亏损则为负数）
流水 Turnover
偿付 Payback

belongs_to :Game
belongs_to :Tenant
*/

package models

import "github.com/jinzhu/gorm"

type RevenueReport struct {
	gorm.Model

	Date     string  `gorm:"type:varchar(10);unique_index:idx_rev_reps"`
	GameID   uint    `gorm:"unique_index:idx_rev_reps;index:idx_game_rev"`
	TenantID uint    `gorm:"unique_index:idx_rev_reps;index:idx_tenant_rev"`
	Payback  float32 `sql:"type:decimal(15, 5)"`
	Turnover uint
}
