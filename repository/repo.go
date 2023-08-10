package repository

import (
	"fmt"

	"github.com/Coreychen4444/Users/model"
	"gorm.io/gorm"
)

type DbRepository struct {
	db *gorm.DB
}

func NewDbRepository(db *gorm.DB) *DbRepository {
	return &DbRepository{db: db}
}

// 创建用户
func (r *DbRepository) CreateUsers(user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// 根据用户名获取用户
func (r *DbRepository) GetUserByName(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

// 根据用户id获取用户
func (r *DbRepository) GetUserById(id int64) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

// 根据文件扩展名创建视频记录

func (r *DbRepository) CreateVideo(video *model.Video) error {
	result := r.db.Create(video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
// repository 包中的 GetPublishedVideos 函数，获取已发布的视频
func (r *DbRepository) GetPublishedVideos() ([]*model.Video, error) {
	var videos []*model.Video
	err := r.db.Find(&videos).Error
	if err != nil {
		return nil, fmt.Errorf("获取视频失败: %w", err)
	}
	return videos, nil
}