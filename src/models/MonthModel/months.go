package monthmodel

import (
	"fetchAPI_gofiber/src/config"

	"gorm.io/gorm"
)

type Month struct {
	gorm.Model
	Name string
	Day  uint
}

func SelectAllMonth() []*Month {
	var items []*Month
	config.DB.Find(&items)
	return items
}
