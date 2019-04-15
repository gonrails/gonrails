package models

type Game struct {
	ID     int    `gorm:"primary_key;column:id"`
	Name   string `gorm:"column:name"`
	NameCn string `gorm:"column:name_cn"`
	Logo   string `gorm:"column:logo"`
}

func (Game) TableName() string {
	return "games"
}
