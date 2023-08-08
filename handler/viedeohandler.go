package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Coreychen4444/Users/service"
	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	s *service.VideoService
}

func NewVideoHandler(s *service.VideoService) *VideoHandler {
	return &VideoHandler{s: s}
}

// 处理视频投稿请求
func (h *VideoHandler) PublishVideo(c *gin.Context) {
	// 从请求头中获取鉴权 token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少鉴权 token"})
		return
	}

	// 读取视频数据
	videoData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "读取视频数据失败"})
		return
	}

	// 获取视频标题并进行非空校验
	title := c.PostForm("title")
	title="title"
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "视频标题不能为空"})
		return
	}

	// 调用 service 层的方法进行视频投稿
	err = h.s.PublishVideo(token, videoData, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("视频投稿失败：%s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "视频投稿成功"})
}

func (h *VideoHandler) GetPublishedVideos(c *gin.Context) {
	// 在这里添加验证用户权限的代码，确保只有登录用户可以访问此接口
	// 你可以使用鉴权 token 来进行验证

	// 调用 service 层的方法获取已发布的视频列表
	videos, err := h.s.GetPublishedVideos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("获取视频列表失败：%s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, videos)
}