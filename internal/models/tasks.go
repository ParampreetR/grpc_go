package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	IsDone      bool
}
