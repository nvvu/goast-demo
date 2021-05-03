package handlers

import (
	"context"
	"fmt"
)

type createUserRequest struct {
	Username string `json:"username"`
}

type createUserResponse struct {
	Message string `json:"message"`
}

// +hdl
func CreateUser(ctx context.Context, req *createUserRequest) (*createUserResponse, error) {
	return &createUserResponse{
		Message: fmt.Sprintf("Hello %s", req.Username),
	}, nil
}

type IDRequest struct {
	ID int64 `json:"id,omitempty"`
}

type UserResponse struct {
	Username string `json:"username,omitempty"`
}

// +hdl
func GetUser(ctx context.Context, req *IDRequest) (*UserResponse, error) {
	return &UserResponse{
		Username: "123",
	}, nil
}
