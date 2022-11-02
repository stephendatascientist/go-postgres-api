package models

import "gorm.io/gorm"

type Employee struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `json:"name"`
	Email   *string `json:"email"`
	Age     uint8   `json:"age"`
	Updated int64   `gorm:"autoUpdateTime:milli" json:"updated"` // Use unix milli seconds as updating time
	Created int64   `gorm:"autoCreateTime" json:"created"`       // Use unix seconds as creating time
}

func MigrateEmployee(db *gorm.DB) error {
	err := db.AutoMigrate(&Employee{})
	return err
}
