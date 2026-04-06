package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequestDTO struct {
	FullName 	string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponseDTO struct {
	ID 			int     `json:"id"`
	Version 	int		`json:"version"`
	FullName 	string	`json:"full_name"`
	PhoneNumber *string	`json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		
	}
}