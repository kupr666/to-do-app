package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/kupr666/to-do-app/internal/core/logger"
	core_http_request "github.com/kupr666/to-do-app/internal/core/transport/http/request"
	core_http_response "github.com/kupr666/to-do-app/internal/core/transport/http/response"
	core_http_types "github.com/kupr666/to-do-app/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title 		core_http_types.Nullable[string]	`json:"title"`
	Description core_http_types.Nullable[string]	`json:"description"`
	Completed 	core_http_types.Nullable[bool] 		`json:"completed"`
}


func (r *PatchTaskRequest) Validate() error {

}

func (h *TasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get taskID")
		return
	}

	var taskRequestDTO PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(
		r,
		&taskRequestDTO,
	); err !=nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
	}
}