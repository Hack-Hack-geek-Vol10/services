package services

import (
	"context"

	user "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/user-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
	storages "github.com/Hack-Hack-geek-Vol10/services/src/storages"
)

type userService struct {
	user.UnimplementedUserServiceServer
	userRepo storages.UserRepo
}

func NewUserService(userRepo storages.UserRepo) user.UserServiceServer {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, arg *user.CreateUserParams) (*user.UserDetail, error) {
	result, err := s.userRepo.Create(ctx, domain.CreateUserParams{
		UserID: arg.UserId,
		Name:   arg.Name,
		Email:  arg.Email,
	})
	if err != nil {
		return nil, err
	}

	return &user.UserDetail{
		UserId: result.UserID,
		Name:   result.Name,
		Email:  result.Email,
	}, nil
}

func (s *userService) GetUser(ctx context.Context, arg *user.GetUserParams) (*user.UserDetail, error) {
	userInfo, err := s.userRepo.ReadOne(ctx, arg.UserId)
	if err != nil {
		return nil, err
	}

	return &user.UserDetail{
		UserId: userInfo.UserID,
		Name:   userInfo.Name,
		Email:  userInfo.Email,
	}, nil
}

func (s *userService) mustEmbedUnimplementedUserServiceServer() {}
