package tasks_http_service

import (
	"context"
	"fmt"

	"github.com/kupr666/to-do-app/internal/core/domain"
)

func (s *TasksService) GetTask(
	ctx context.Context,
	userID int,
) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, userID)
	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"failed to get task from repository: %w",
			err,
		)
	}

	return task, nil
}