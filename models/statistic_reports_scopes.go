package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

func (v *StatisticReports) withDate(date string) *gorm.DB {
	return db.Model(v).Where(&StatisticReports{Date: date})
}

func (v *StatisticReports) withGame(GameID uint) *gorm.DB {
	return db.Model(v).Where(&StatisticReports{GameID: GameID})
}

func (v *StatisticReports) withTenant(TenantID uint) *gorm.DB {
	return db.Model(v).Where(&StatisticReports{TenantID: TenantID})
}

func (v *StatisticReports) FilterParams(params map[string]string) *gorm.DB {

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
