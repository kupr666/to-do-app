package tasks_http_service

import (
	"context"
	"fmt"

	"github.com/kupr666/to-do-app/internal/core/domain"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	taskID int,
	patch domain.TaskPatch,
) (domain.Task, error) {

	task, err := s.tasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"failed to get task from repository: %w",
			err,
		)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("applly taks patch: %w", err)
	}

	patchedTask, err := s.tasksRepository.PatchTask(ctx, taskID, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	return patchedTask, nil
}
