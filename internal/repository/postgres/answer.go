package postgres

import (
	"context"
	"fmt"

	"github.com/tmozzze/QuesAns/internal/model"
	"gorm.io/gorm"
)

type AnswerRepo struct {
	db *gorm.DB
}

func NewAnswerRepo(db *gorm.DB) *AnswerRepo {
	return &AnswerRepo{db: db}
}

func (r *AnswerRepo) Create(ctx context.Context, a *model.Answer) error {
	const op = "repository.postgres.answer.AnswerRepo.Create"

	err := r.db.WithContext(ctx).Create(a).Error
	if err != nil {
		return fmt.Errorf("%s: failed to create answer: %w", op, err)
	}

	return nil
}

func (r *AnswerRepo) GetById(ctx context.Context, id uint) (*model.Answer, error) {
	const op = "repository.postgres.answer.AnswerRepo.GetById"

	var a model.Answer
	err := r.db.WithContext(ctx).First(&a, id).Error
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get answer by id: %v: %w", op, id, err)
	}
	return &a, nil
}

func (r *AnswerRepo) Delete(ctx context.Context, id uint) error {
	const op = "repository.postgres.answer.AnswerRepo.Delete"

	err := r.db.WithContext(ctx).Delete(&model.Answer{}, id).Error
	if err != nil {
		return fmt.Errorf("%s: failed to delete answer by id: %v: %w", op, id, err)
	}
	return nil
}
