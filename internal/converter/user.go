package converter

import (
	"github.com/Danya97i/auth/internal/models"
	pb "github.com/Danya97i/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserFromService конвертирует models.User в pb.User
func ToUserFromService(user models.User) *pb.User {
	pbUser := &pb.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
	if user.UpdatedAt != nil {
		pbUser.UpdatedAt = timestamppb.New(*user.UpdatedAt)
	}
	return pbUser
}

// ToUserInfoFromService конвертирует models.UserInfo в pb.UserInfo
func ToUserInfoFromService(userInfo models.UserInfo) *pb.UserInfo {
	return &pb.UserInfo{
		Name:  userInfo.Name,
		Email: userInfo.Email,
		Role:  ToRoleFromService(userInfo.Role),
	}
}

// ToRoleFromService конвертирует pb.Role в models.Role
func ToRoleFromService(r models.Role) pb.Role {
	switch r {
	case models.ADMIN:
		return pb.Role_ADMIN
	case models.USER:
		return pb.Role_USER
	default:
		return pb.Role_UNKNOWN
	}
}

// --------

// ToUserInfoFromPb конвертирует pb.UserInfo в models.UserInfo
func ToUserInfoFromPb(pbInfo *pb.UserInfo) *models.UserInfo {
	if pbInfo == nil {
		return nil
	}
	info := models.UserInfo{
		Name:  pbInfo.Name,
		Email: pbInfo.Email,
		Role:  ToRoleFromPb(pbInfo.Role),
	}
	return &info
}

// ToRoleFromPb конвертирует pb.Role в models.Role
func ToRoleFromPb(r pb.Role) models.Role {
	switch r {
	case pb.Role_ADMIN:
		return models.ADMIN
	case pb.Role_USER:
		return models.USER
	default:
		return models.UNKNOWN
	}
}
