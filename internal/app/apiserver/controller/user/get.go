package user

import "github.com/gin-gonic/gin"

func (u *Ctrl) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get one user",
	})
}
