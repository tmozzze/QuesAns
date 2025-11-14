package postgres

import "gorm.io/gorm"

type Repository struct {
	Question *QuestionRepo
	Answer   *AnswerRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Question: NewQuestionRepo(db),
		Answer:   NewAnswerRepo(db),
	}
}
