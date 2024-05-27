package apiserver

import (
	"github.com/gin-gonic/gin"
	"go-auth/internal/app/apiserver/controller/user"
)

func initRouter(g *gin.Engine) {
	usr := g.Group("/users")
	{
		uc := user.NewCtrl()

		usr.GET("", uc.List)
		usr.GET("/:id", uc.Get)
		usr.PUT("/:id", uc.Update)
	}
}
