/*
Author: 曾涛
Desc:   重新设计统计模块 南朝四百八十寺，多少楼台烟雨中
Date:   2019-05-06
Email:  zengtao@risewinter.com
*/

/*
User Story
商户
游戏
日期
营收（如果亏损则为负数）
流水 Turnover
偿付 Payback
系列赛数量（开盘）
盘口数量
竞猜用户数量
新增竞猜用户数量
有效订单数量

belongs_to :Game
belongs_to :Tenant

UNIQ Index: (Date, Tenant, Game, TopicType)
*/

package models

import "github.com/jinzhu/gorm"

type StatisticReports struct {
	gorm.Model

	Date              string  `gorm:"type:varchar(10);unique_index:idx_rev_reps"`
	GameID            uint    `gorm:"unique_index:idx_rev_reps;index:idx_game_rev"`
	TenantID          uint    `gorm:"unique_index:idx_rev_reps;index:idx_tenant_rev"`
	TopicType         uint    `gorm:"unique_index:idx_rev_reps"`
	Payback           float32 `sql:"type:decimal(15, 5)"`
	Turnover          uint    `gorm:"default:0"`
	SeriesCount       uint    `gorm:"default:0"`
	TopicCount        uint    `gorm:"default:0"`
	BetUserCount      uint    `gorm:"default:0"`
	FreshBetUserCount uint    `gorm:"default:0"`
	BetOrderCount     uint    `gorm:"default:0"`
}

func (StatisticReports) TableName() string {
	return "statistic_reports"
}
