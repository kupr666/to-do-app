package tasks_postgres_repository

import (
	"time"

	"github.com/kupr666/to-do-app/internal/core/domain"
)

type TaskModel struct {
	ID 			 int
	Version 	 int
	Title 		 string
	Description  *string
	Completed 	 bool
	CreatedAt 	 time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}


func taskDomainsFromModels(tasks []TaskModel) []domain.Task {
	tasksDomains := make([]domain.Task, len(tasks))

	for i, task := range tasks {
		tasksDomains[i] = domain.NewTask(
			task.ID,
			task.Version,
			task.Title,
			task.Description,
			task.Completed,
			task.CreatedAt,
			task.CompletedAt,
			task.AuthorUserID,
		)
	}

	return tasksDomains
}