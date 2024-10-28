package user

import (
	"context"
	"errors"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Danya97i/auth/internal/converter"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/service"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// Server - имплементация gRPC-сервера
type Server struct {
	pb.UnimplementedUserV1Server
	userService service.UserService
}

// NewServer - конструктор сервера
func NewServer(userService service.UserService) *Server {
	return &Server{userService: userService}
}

// CreateUser - метод для создания нового пользователя
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("create user request: ", req)

	info := converter.ToUserInfoFromPb(req.Info)
	if info == nil {
		return nil, errors.New("invalid info")
	}
	id, err := s.userService.CreateUser(ctx, *info, req.Password, req.PasswordConfirm)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Id: id}, nil
}

// GetUser - метод для получения информации о пользователе
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("get user request: ", req)

	user, err := s.userService.User(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: converter.ToUserFromService(*user),
	}, nil
}

// UpdateUser - метод обновления данных пользователя
func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Println("update user request: ", req)
	userInfo := models.UserInfo{
		Role: converter.ToRoleFromPb(req.Role),
	}

	if req.Name != nil {
		userInfo.Name = req.Name.Value
	}

	if err := s.userService.UpdateUser(ctx, req.Id, &userInfo); err != nil {
		return nil, err
	}

	return nil, nil
}

// DeleteUser - метод удаления пользователя
func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Println("delete user request: ", req)

	if err := s.userService.DeleteUser(ctx, req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}
