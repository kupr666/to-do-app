package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/kupr666/to-do-app/internal/core/domain"
	core_http_server "github.com/kupr666/to-do-app/internal/core/transport/http/server"
)

type StatService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

type StatisticsHTTPHandler struct {
	statService StatService
}

func NewStatHTTPHandler(statService StatService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statService: statService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
