package user

import (
	"github.com/gin-gonic/gin"
	"go-auth/pkg/rdb"
)

func (u *Ctrl) Get(c *gin.Context) {
	get := rdb.Client().Get(c, "a")

	c.JSON(200, gin.H{
		"message": rdb.Client().PoolStats(),
		"data":    get.Val(),
		"err":     get.Val(),
	})
}
