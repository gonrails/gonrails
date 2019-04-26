/*
Author: 曾涛
Desc:   RevenueReport 的一些查询过滤条件
Date:   2019-04-26
Email:  zengtao@risewinter.com
*/

package models

import "github.com/jinzhu/gorm"

func (v *RevenueReport) withDate(date string) *gorm.DB {
	return nil
}

func (v *RevenueReport) withGame(GameID uint) *gorm.DB {
	return nil
}

func (v *RevenueReport) withTenant(TenantID uint) *gorm.DB {
	return nil
}
