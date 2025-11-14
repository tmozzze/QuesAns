package model

import "time"

type Answer struct {
	Id         uint      `gorm:"primaryKey"`
	QuestionId uint      `gorm:"not null;index"`
	UserId     string    `gorm:"type:uuid;not null"`
	Text       string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
