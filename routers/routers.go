package routers

import (
	"github.com/Coreychen4444/Users/handler"
	"github.com/Coreychen4444/Users/repository"
	"github.com/Coreychen4444/Users/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitRouter 初始化路由信息
func InitRouter(db *gorm.DB) *gin.Engine {
// 创建一个默认的gin引擎

router := gin.Default()

// 创建数据库仓库
repo := repository.NewDbRepository(db)

// 创建用户服务
userService := service.NewUserService(repo)

// 创建视频服务
videoService := service.NewVideoService(repo)

// 创建用户处理器
userHandler := handler.NewUserHandler(userService)

// 创建视频处理器
videoHandler := handler.NewVideoHandler(videoService)

// 创建用户组路由
userGroup := router.Group("/douyin/user")
{
    userGroup.POST("/register", userHandler.Register) // 注册
    userGroup.POST("/login", userHandler.Login) // 登录
    userGroup.GET("/", userHandler.GetUserInfo) // 获取用户信息
}

// 创建视频组路由
videoGroup := router.Group("/douyin/video")
{
	videoGroup.POST("/publish", videoHandler.PublishVideo) // 发布视频
}

// 返回路由
return router
}