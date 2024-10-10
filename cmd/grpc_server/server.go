package main

import (
	"context"
	"log"

	pb "github.com/Danya97i/auth/pkg/user_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedUserV1Server
}

// CreateUser - метод для создания нового пользователя
func (s *server) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("create user request: ", req)
	return nil, nil
}

// GetUser- метод для получения информации о пользователе
func (s *server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("get user request: ", req)
	return nil, nil
}

// UpdateUser - метод обновления данных пользователя
func (s *server) UpdateUser(_ context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Println("update user request: ", req)
	return nil, nil
}

// DeleteUser - метод удаления пользователя
func (s *server) DeleteUser(_ context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Println("delete user request: ", req)
	return nil, nil
}
