package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kupr666/to-do-app/internal/core/domain"
	core_logger "github.com/kupr666/to-do-app/internal/core/logger"
	core_http_request "github.com/kupr666/to-do-app/internal/core/transport/http/request"
	core_http_response "github.com/kupr666/to-do-app/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated           int      `json:"tasks_created"`
	TasksCompleted         int      `json:"tasks_completed"`
	TasksCompletedRate     *float64 `json:"tasks_completed_rate"`
	TasksAvgCompletionTime *string  `json:"tasks_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, from, to, err := GetUserIDFromTo(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'user_id'/'from'/'to' query param",
		)
		return
	}

	statistics, err := h.statService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistics")
		return
	}

	response := DTOFromDomain(statistics)

	responseHandler.JsonResponse(response, http.StatusOK)
}

func DTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TasksAvgCompletionTime != nil {
		duration := statistics.TasksAvgCompletionTime.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TasksCreated:           statistics.TasksCreated,
		TasksCompleted:         statistics.TasksCompleted,
		TasksCompletedRate:     statistics.TasksCompletedRate,
		TasksAvgCompletionTime: avgTime,
	}
}

func GetUserIDFromTo(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParam = "user_id"
		fromQueryParam   = "from"
		toQueryParam     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'userID' Query param: %w", err)
	}

	from, err := core_http_request.GetTimeQueryParam(r, fromQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' Query apram: %w", err)
	}

	to, err := core_http_request.GetTimeQueryParam(r, toQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' Query param: %w", err)
	}

	return userID, from, to, nil
}
