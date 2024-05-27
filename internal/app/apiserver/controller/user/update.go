package user

import "github.com/gin-gonic/gin"

func (u *Ctrl) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update user",
	})
}
