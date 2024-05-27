package repo

import (
	"go-auth/internal/app/pkg/apiserver/model"
)

type UserRepo interface {
	List() model.UserList
}
