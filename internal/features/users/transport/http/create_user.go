package users_transport_http

import (
	"net/http"

	"github.com/kupr666/to-do-app/internal/core/domain"
	core_logger "github.com/kupr666/to-do-app/internal/core/logger"
	core_http_request "github.com/kupr666/to-do-app/internal/core/transport/http/request"
	core_http_response "github.com/kupr666/to-do-app/internal/core/transport/http/response"
)

type CreateUserRequestDTO struct {
	FullName 	string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponseDTO UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var user CreateUserRequestDTO
	if err := core_http_request.DecodeAndValidateRequest(r, &user); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(user)
	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponseDTO(userDTOFromDomain(userDomain))

	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequestDTO) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
