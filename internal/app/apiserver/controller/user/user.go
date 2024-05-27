package user

import "go-auth/internal/app/apiserver/service"

type Ctrl struct {
	userSrv service.UserService
}

func NewCtrl() *Ctrl {
	return &Ctrl{
		userSrv: service.NewUserService(),
	}
}
