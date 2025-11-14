package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/tmozzze/QuesAns/internal/model"
)

// Repo interface
type AnswerRepo interface {
	Create(ctx context.Context, a *model.Answer) error
	GetById(ctx context.Context, id uint) (*model.Answer, error)
	Delete(ctx context.Context, id uint) error
}

type AnswerService struct {
	repo AnswerRepo
	log  *slog.Logger
}

func NewAnswerService(r AnswerRepo, log *slog.Logger) *AnswerService {
	return &AnswerService{repo: r, log: log}
}

// Business

// Create - Create answer
func (s *AnswerService) Create(ctx context.Context, questionId uint, userId, text string) (*model.Answer, error) {
	const op = "service.AnswerService.Create"
	svcLog := s.log.With(slog.String("op", op))

	svcLog.Debug("Create called", "questionId", questionId, "userId", userId, "text", text)

	if text == "" {
		svcLog.Info("attempt to create empty question")
		return nil, fmt.Errorf("%s: text cannot be empty", op)
	}

	a := &model.Answer{
		QuestionId: questionId,
		UserId:     userId,
		Text:       text,
	}
	if err := s.repo.Create(ctx, a); err != nil {
		svcLog.Error("failed to create answer in repo", "err", err)
		return nil, fmt.Errorf("%s: failed to create answer: %w", op, err)
	}

	svcLog.Info("answer created", "questionId", questionId)
	svcLog.Debug("answer created", "answer", a)

	return a, nil

}

// GetById - Get answer by id
func (s *AnswerService) GetById(ctx context.Context, id uint) (*model.Answer, error) {
	const op = "service.AnswerService.GetById"
	svcLog := s.log.With(slog.String("op", op))

	svcLog.Debug("GetById called", "id", id)

	a, err := s.repo.GetById(ctx, id)
	if err != nil {
		svcLog.Error("failed to get answer by id in repo", "err", err)
		return nil, fmt.Errorf("%s: failed to get answer by id: %w", op, err)
	}

	return a, nil

}

// Delete - Delete answer
func (s *AnswerService) Delete(ctx context.Context, id uint) error {
	const op = "service.AnswerService.Delete"
	svcLog := s.log.With(slog.String("op", op))

	svcLog.Debug("Delete called", "id", id)

	if err := s.repo.Delete(ctx, id); err != nil {
		svcLog.Error("failed to delete answer in repo", "err", err)
		return fmt.Errorf("%s: failed to delete answer %d: %w", op, id, err)
	}

	svcLog.Info("answer deleted", "id", id)

	return nil
}
