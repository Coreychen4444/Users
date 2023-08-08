package model

import "time"

// User
type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
	Username        string `json:"-" gorm:"unique"`  // 注册用户名，最长32个字符
	PasswordHash    string `json:"-" gorm:"size:100"`                // 密码，最长32个字符   service层完成对应的逻辑操作
}
type Video struct {
	ID            uint      `gorm:"primaryKey" json:"id"` // 视频ID
	AuthorID      int64     `json:"author_id"` // 作者ID
	Title         string    `json:"title"` // 视频标题
	VideoFilePath string    `json:"video_file_path"` // 视频文件路径
	CreateTime    time.Time `json:"create_time"` // 创建时间
}