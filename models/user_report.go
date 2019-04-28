/*
Author: 曾涛
Desc:   UserReport 用户数据报表
Date:   2019-04-28
Email:  zengtao@risewinter.com
*/

// 新增下注用户数（）

package models

import "github.com/jinzhu/gorm"

type UserReport struct {
	gorm.Model
	Date       string `gorm:"type:varchar(10);unique_index:idx_rev_reps"`
	GameID     uint   `gorm:"unique_index:idx_rev_reps;index:idx_game_rev"`
	TenantID   uint   `gorm:"unique_index:idx_rev_reps;index:idx_tenant_rev"`
	BetCount   uint
	FreshCount uint
	OrderCount uint
}
