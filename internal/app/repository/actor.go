package repository

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	ID   uint64 `json:"ID" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
