package user

import (
	"context"
	"github.com/gin-gonic/gin"
)

// List 获取用户列表
func (u *Ctrl) List(c *gin.Context) {
	list, err := u.userSrv.List(context.Background())
	if err != nil {
		panic(err)
	}

	c.JSON(200, list)
}
