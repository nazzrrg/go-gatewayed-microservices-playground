package service

import (
	"context"
	"log"
)

type UserService struct{}

func (s *UserService) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) CreateUser(ctx context.Context, m *CreateUserRequest) (*Response, error) {
	log.Printf("Create user: %s", m.String())
	return &Response{Success: true}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, m *UpdateUserRequest) (*Response, error) {
	log.Printf("Update user: %s", m.String())
	return &Response{Success: true}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, m *DeleteUserRequest) (*Response, error) {
	log.Printf("Delete user: %s", m.String())
	return &Response{Success: true}, nil
}
