package mysql

import (
	"go-auth/internal/app/apiserver/repo"
	"go-auth/internal/app/pkg/apiserver/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

var _ repo.UserRepo = (*userRepo)(nil)

func newUserRepo(ds *dataset) *userRepo {
	return &userRepo{db: ds.db}
}

func (r userRepo) List() model.UserList {
	var users []*model.User
	r.db.Find(&users)
	return model.UserList{Items: users}
}

// 确保userRepo实现了 UserRepo 接口
