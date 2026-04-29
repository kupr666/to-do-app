package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/kupr666/to-do-app/internal/core/logger"
	core_http_response "github.com/kupr666/to-do-app/internal/core/transport/http/response"
	core_http_request "github.com/kupr666/to-do-app/internal/core/transport/http/request"
)

type GetTasksResponseDTO []TaskDTOResponse

func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, limit, offset, err := GetUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'userID'/'limit'/'offset' query param",
		)
		return
	}

	tasksDomains, err := h.tasksService.GetTasks(
		ctx,
		userID,
		limit,
		offset,
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}

	response := GetTasksResponseDTO(tasksDTOFromDomain(tasksDomains))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func GetUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		userIDQueryParamKey = "user_id"
		limitQueryParamKey = "limit"
		offsetQueryParamKey = "offset"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'userID' Query param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return userID, limit, offset, nil
}
