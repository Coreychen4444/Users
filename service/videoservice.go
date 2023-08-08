package service

import (
	"errors"
	"fmt"
	"github.com/Coreychen4444/Users/model"
	"github.com/Coreychen4444/Users/repository"
	// "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type VideoService struct {


	r *repository.DbRepository // DbRepository类型的指针r
}
func NewVideoService(r *repository.DbRepository) *VideoService {


	return &VideoService{r: r}
}
// 根据文件扩展名创建视频服务
// 参数：r 数据库仓库
// 返回值：视频服务指针

// 视频投稿
func (s *VideoService) PublishVideo(token string, videoData []byte, title string) error {
	claims, err := VerifyToken(token)
	if err != nil {
		return fmt.Errorf("验证token时出错: %w", err)
	}

	userID := claims.UserID

	user, err := s.r.GetUserById(userID)
	if err != nil {
		return fmt.Errorf("获取用户信息时出错: %w", err)
	}

	if user == nil {
		return errors.New("用户不存在")
	}

	// 处理视频数据，保存到本地或对象存储，这里假设保存在本地
	videoPath := fmt.Sprintf("videos/%d_%s.mp4", userID, time.Now().Format("20060102150405"))
	if err := s.saveVideoData(videoData, videoPath); err != nil {
		return fmt.Errorf("保存视频文件时出错: %w", err)
	}

	// 创建视频记录
	video := &model.Video{
		AuthorID:      userID,
		Title:         title,
		VideoFilePath: videoPath,
		CreateTime:    time.Now(),
	}

	if err := s.r.CreateVideo(video); err != nil {
		return fmt.Errorf("创建视频记录时出错: %w", err)
	}

	return nil
}

// 保存视频数据
func (s *VideoService) saveVideoData(data []byte, filePath string) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (s *VideoService) GetPublishedVideos() ([]*model.Video, error) {

	// 从仓库获取已发布视频
	videos, err := s.r.GetPublishedVideos()
	if err != nil {
		// 如果获取视频失败，则返回错误信息
		return nil, fmt.Errorf("从仓库获取视频失败: %w", err)
	}
	// 返回已发布视频
	return videos, nil
}