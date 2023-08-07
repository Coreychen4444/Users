package routers

import (
	"github.com/Coreychen4444/Users/handler"
	"github.com/Coreychen4444/Users/repository"
	"github.com/Coreychen4444/Users/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitRouter initialize routing information
func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	repo := repository.NewDbRepository(db)
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)
	user := r.Group("/douyin/user")
	{
		user.POST("/register", userHandler.Register)
		user.POST("/login", userHandler.Login)
		user.GET("/", userHandler.GetUserInfo)
	}
	return r
}
