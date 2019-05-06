/*
Author: 曾涛
Desc:   RevenueReport 的一些查询过滤条件
Date:   2019-04-26
Email:  zengtao@risewinter.com
*/

package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

// 到时候这里可以抽象成为通用的 v interface{}
func (v *RevenueReport) withDate(date string) *gorm.DB {
	return db.Model(v).Where(&RevenueReport{Date: date})
}

func (v *RevenueReport) withGame(GameID uint) *gorm.DB {
	return db.Model(v).Where(&RevenueReport{GameID: GameID})
}

func (v *RevenueReport) withTenant(TenantID uint) *gorm.DB {
	return db.Model(v).Where(&RevenueReport{TenantID: TenantID})
}

func (v *RevenueReport) FilterParams(params map[string]string) *gorm.DB {

	var records = db.Model(v)

	if value, ok := params["tenant"]; ok {
		id, _ := strconv.Atoi(value)
		var uid uint = uint(id)
		records = v.withTenant(uid)
	}

	if value, ok := params["game"]; ok {
		id, _ := strconv.Atoi(value)
		var uid uint = uint(id)
		records = v.withGame(uid)
	}

	return records
}
