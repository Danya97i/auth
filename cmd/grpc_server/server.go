package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// Role - роль пользователя
type Role int32

// Значения роли
const (
	UNKNOWN Role = iota
	ADMIN
	USER
)

func (r Role) String() string {
	switch r {
	case ADMIN:
		return "admin"
	case USER:
		return "user"
	default:
		return "unknown"
	}
}

func parseRole(role string) Role {
	switch role {
	case ADMIN.String():
		return ADMIN
	case USER.String():
		return USER
	default:
		return UNKNOWN
	}
}

type server struct {
	pb.UnimplementedUserV1Server
	pool *pgxpool.Pool
}

// CreateUser - метод для создания нового пользователя
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("create user request: ", req)

	// хеширование пароля
	pswrd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// создание пользователя в БД
	query := "INSERT INTO users (name, email, password, role, created_at) VALUES ($1, $2, $3, $4, $5) returning id;"
	row := s.pool.QueryRow(ctx, query, req.Name, req.Email, string(pswrd), parseRoleFromPb(req.Role).String(), time.Now())

	var id int64
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	log.Println("user id: ", id)
	return &pb.CreateUserResponse{Id: id}, nil
}

// GetUser- метод для получения информации о пользователе
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("get user request: ", req)

	query := "SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1;"

	var (
		id                   int32
		name                 string
		email                string
		role                 string
		createdAt, updatedAt sql.NullTime
	)
	row := s.pool.QueryRow(ctx, query, req.Id)
	if err := row.Scan(&id, &name, &email, &role, &createdAt, &updatedAt); err != nil {
		return nil, err
	}

	log.Println("user id: ", id, "name: ", name, "email: ",
		email, "role: ", role, "created_at: ", createdAt, "updated_at: ", updatedAt,
	)

	return &pb.GetUserResponse{
		Id:        int64(id),
		Name:      name,
		Email:     email,
		Role:      pbRole(parseRole(role)),
		CreatedAt: timestamppb.New(createdAt.Time),
		UpdatedAt: timestamppb.New(updatedAt.Time),
	}, nil
}

// UpdateUser - метод обновления данных пользователя
func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Println("update user request: ", req)

	queryBuilder := squirrel.Update("users").
		PlaceholderFormat(squirrel.Dollar).
		Set("role", parseRoleFromPb(req.Role).String()).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": req.Id})

	if req.Name != nil {
		queryBuilder = queryBuilder.Set("name", req.Name.GetValue())
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := s.pool.Exec(ctx, query, args...); err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteUser - метод удаления пользователя
func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Println("delete user request: ", req)

	queryBuilder := squirrel.Delete("users").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": req.Id})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	if _, err := s.pool.Exec(ctx, query, args...); err != nil {
		return nil, err
	}
	return nil, nil
}

func pbRole(role Role) pb.Role {
	switch role {
	case ADMIN:
		return pb.Role_ADMIN
	case USER:
		return pb.Role_USER
	default:
		return pb.Role_UNKNOWN
	}
}

func parseRoleFromPb(pbRole pb.Role) Role {
	switch pbRole {
	case pb.Role_ADMIN:
		return ADMIN
	case pb.Role_USER:
		return USER
	default:
		return UNKNOWN
	}
}
