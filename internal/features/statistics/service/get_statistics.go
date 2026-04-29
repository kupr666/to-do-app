package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/kupr666/to-do-app/internal/core/domain"
	core_errors "github.com/kupr666/to-do-app/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"'to' must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.statRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"failed to get tasks from repository: %w",
			err,
		)
	}

	statistics := domain.CreateStatistics(tasks)

	return statistics, nil
}
