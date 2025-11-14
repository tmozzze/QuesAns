package model

import "time"

type Question struct {
	Id        uint      `gorm:"primaryKey"`
	Text      string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Answers   []Answer  `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}
