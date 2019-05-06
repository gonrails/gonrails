package models

import (
	"github.com/jinzhu/gorm"
)

type TopicReport struct {
	gorm.Model
	Date     string `gorm:"type:varchar(10);unique_index:idx_rev_reps"`
	GameID   uint   `gorm:"unique_index:idx_rev_reps;index:idx_game_rev"`
	TenantID uint   `gorm:"unique_index:idx_rev_reps;index:idx_tenant_rev"`
}
