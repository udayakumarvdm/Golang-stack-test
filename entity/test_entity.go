package model

import "time"

type TestEntity struct {
	StartDate time.Time `gorm:"column:start_date"`
	EndDate   time.Time `gorm:"column:end_date"`
}

func (TestEntity) TableName() string {
	return "dates"
}
