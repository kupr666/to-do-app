package tasks_transport_http

import (
	"time"
	"net/http"

	"github.com/kupr666/to-do-app/internal/core/domain"
	core_logger "github.com/kupr666/to-do-app/internal/core/logger"
	core_http_request "github.com/kupr666/to-do-app/internal/core/transport/http/request"
	core_http_response "github.com/kupr666/to-do-app/internal/core/transport/http/response"
)

type CreateTaskRequestDTO struct {
	Title 		 string  `json:"title" validate:"required,min=1,max=100"`
	Describtion  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserID int 	 `json:"author_user_id" validate:"required"`
}

type CreateTaskResponseDTO struct {
	ID 			 int 		`json:"id"`
	Version 	 int 		`json:"version"`
	Title 		 string 	`json:"title"`
	Description  *string 	`json:"description"`
	Completed 	 bool 		`json:"completed"`
	CreatedAt 	 time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int 		`json:"author_user_id"`
}

func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) () {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateTaskRequestDTO

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	taskDomain := domain.NewTaskUninitialized(
		request.Title,
		request.Describtion,
		request.AuthorUserID,
	)

	taskDomain, err := h.tasksService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "filed to create task")
		return
	}

	response := taskDTOFromDomain(taskDomain)

	responseHandler.JsonResponse(response, http.StatusCreated)
}

func taskDTOFromDomain(task domain.Task) CreateTaskResponseDTO {
	return CreateTaskResponseDTO{
		ID: task.ID,
		Version: task.Version,
		Title: task.Title,
		Description: task.Description,
		Completed: task.Completed,
		CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}
