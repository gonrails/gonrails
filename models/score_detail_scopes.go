/*
Author: 曾涛
Desc:   ScoreDetail 的一些查询过滤条件
Date:   2019-04-26
Email:  zengtao@risewinter.com
*/

package models

import "github.com/jinzhu/gorm"

func (v *ScoreDetail) withDate(date string) *gorm.DB {
	return db.Model(v).Where(&ScoreDetail{Date: date})
}

func (v *ScoreDetail) withGame(gameID uint) *gorm.DB {
	return db.Model(v).Where(&ScoreDetail{GameID: gameID})
}

func (v *ScoreDetail) withTenant(tenantID uint) *gorm.DB {
	return db.Model(v).Where(&ScoreDetail{TenantID: tenantID})
}

// 验证

// Sidekiq 停掉
// 另外一台服务器上起 Sidekiq
