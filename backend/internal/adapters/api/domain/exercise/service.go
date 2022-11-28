package exercise

import (
	"context"
	"polyfit/pkg/logging"
)

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]Exercise, error) {
	return s.repository.FindAll(ctx)
}
