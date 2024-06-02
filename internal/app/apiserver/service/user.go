package service

import (
	"context"
	"go-auth/internal/app/apiserver/repo"
	"go-auth/internal/app/pkg/apiserver/model"
)

type UserService interface {
	List(ctx context.Context) (*model.UserList, error)
}

type userService struct {
}

var _ UserService = (*userService)(nil)

func NewUserService() UserService {
	return &userService{}
}

// List 返回用户列表
func (u userService) List(ctx context.Context) (*model.UserList, error) {
	list := repo.Client().User().List()
	return &list, nil
}
