package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/tmozzze/QuesAns/internal/model"
)

// Repo interface
type QuestionRepo interface {
	Create(ctx context.Context, q *model.Question) error
	GetById(ctx context.Context, id uint) (*model.Question, error)
	List(ctx context.Context, limit, offset int) ([]model.Question, error)
	Delete(ctx context.Context, id uint) error
}

type QuestionService struct {
	repo QuestionRepo
	log  *slog.Logger
}

func NewQuestionService(r QuestionRepo, log *slog.Logger) *QuestionService {
	return &QuestionService{repo: r, log: log}
}

// Business

// Create - Create question
func (s *QuestionService) Create(ctx context.Context, text string) (*model.Question, error) {
	const op = "service.QuestionService.Create"
	svcLog := s.log.With(slog.String("op", op))

	svcLog.Debug("Create called", "text", text)

	if text == "" {
		svcLog.Info("attempt to create empty question")
		return nil, fmt.Errorf("%s: text cannot be empty", op)
	}

	q := &model.Question{Text: text}

	if err := s.repo.Create(ctx, q); err != nil {
		svcLog.Error("failed to create question in repo", "err", err)
		return nil, fmt.Errorf("%s: failed to create question: %w", op, err)
	}

	svcLog.Info("question created", "id", q.Id)
	svcLog.Debug("question created", "questions", q)

	return q, nil
}

// GetById - Get question by id
func (s *QuestionService) GetById(ctx context.Context, id uint) (*model.Question, error) {
	const op = "service.QuestionService.GetById"
	svcLog := s.log.With(slog.String("op", op))

	svcLog.Debug("GetById called", "id", id)

	q, err := s.repo.GetById(ctx, id)
	if err != nil {
		svcLog.Error("failed to get question by id in repo", "err", err)
		return nil, fmt.Errorf("%s: failed to get question by id %d: %w", op, id, err)
	}

	return q, nil
}

// List - Get list of questions
func (s *QuestionService) List(ctx context.Context, limit, offset int) ([]model.Question, error) {
	const op = "service.QuestionService.List"
	svcLog := s.log.With(slog.String("op", op))

	svcLog.Debug("List called", "limit", limit, "offset", offset)

	list, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		svcLog.Error("failed to get list of questions in repo", "err", err)
		return nil, fmt.Errorf("%s: failed to list questions: %w", op, err)
	}

	return list, nil
}

// Delete - Delete question by id
func (s *QuestionService) Delete(ctx context.Context, id uint) error {
	const op = "service.QuestionService.Delete"
	svcLog := s.log.With(slog.String("op", op))

	svcLog.Debug("Delete called", "id", id)

	if err := s.repo.Delete(ctx, id); err != nil {
		svcLog.Error("failed to delete question in repo", "err", err)
		return fmt.Errorf("%s: failed to delete question %d: %w", op, id, err)
	}

	svcLog.Info("question deleted", "id", id)

	return nil
}
