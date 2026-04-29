package statistics_service

import (
	"context"
	"time"

	"github.com/kupr666/to-do-app/internal/core/domain"
)

type StatisticsService struct {
	statRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatService(repo StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		statRepository: repo,
	}
}
