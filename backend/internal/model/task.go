package model

import "time"

type Task struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:text;not null"`
	Description *string   `json:"description,omitempty" gorm:"type:text"`
	Status      string    `json:"status" gorm:"type:varchar(50);default:'pending'"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
