package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Danya97i/auth/internal/models"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// ToUserFromService конвертирует models.User в pb.User
func ToUserFromService(user *models.User) *pb.User {
	if user == nil {
		return nil
	}

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
func ToUserInfoFromService(userInfo *models.UserInfo) *pb.UserInfo {
	if userInfo == nil {
		return nil
	}

	pbUserInfo := pb.UserInfo{
		Email: userInfo.Email,
		Role:  ToRoleFromService(userInfo.Role),
	}

	if userInfo.Name != nil {
		pbUserInfo.Name = *userInfo.Name

	}

	return &pbUserInfo
}

// --------

// ToUserInfoFromPb конвертирует pb.UserInfo в models.UserInfo
func ToUserInfoFromPb(pbInfo *pb.UserInfo) *models.UserInfo {
	if pbInfo == nil {
		return nil
	}

	info := models.UserInfo{
		Name:  &pbInfo.Name,
		Email: pbInfo.Email,
		Role:  ToRoleFromPb(pbInfo.Role),
	}

	return &info
}

// ToUserInfoFromPbUpdateRequest собирает UserInfo из UpdateUserRequest
func ToUserInfoFromPbUpdateRequest(r *pb.UpdateUserRequest) *models.UserInfo {
	userInfo := models.UserInfo{
		Role: ToRoleFromPb(r.Role),
	}

	if r.Name != nil {
		userInfo.Name = &r.Name.Value
	}

	return &userInfo
}
