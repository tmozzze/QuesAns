package postgres

import (
	"context"
	"fmt"

	"github.com/tmozzze/QuesAns/internal/model"
	"gorm.io/gorm"
)

type QuestionRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{db: db}
}

func (r *QuestionRepo) Create(ctx context.Context, q *model.Question) error {
	const op = "repository.question.postgres.QuestionRepo.Create"

	err := r.db.WithContext(ctx).Create(q).Error
	if err != nil {
		return fmt.Errorf("%s: failed to create question: %w", op, err)
	}

	return nil
}

func (r *QuestionRepo) GetById(ctx context.Context, id uint) (*model.Question, error) {
	const op = "repository.question.postgres.QuestionRepo.GetById"

	var q model.Question
	err := r.db.WithContext(ctx).First(&q, id).Error
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get question by id: %v: %w", op, id, err)
	}
	return &q, nil
}

func (r *QuestionRepo) List(ctx context.Context, limit, offset int) ([]model.Question, error) {
	const op = "repository.question.postgres.QuestionRepo.List"

	var list []model.Question
	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&list).Error

	if err != nil {
		return nil, fmt.Errorf("%s: failed to get list of questions (limit=%d offset=%d): %w", op, limit, offset, err)
	}

	return list, nil
}

func (r *QuestionRepo) Delete(ctx context.Context, id uint) error {
	const op = "repository.question.postgres.QuestionRepo.Delete"

	err := r.db.WithContext(ctx).Delete(&model.Question{}, id).Error
	if err != nil {
		return fmt.Errorf("%s: failed to delete question by id: %v: %w", op, id, err)
	}
	return nil
}
