// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID        int32     `gorm:"column:user_id;not null" json:"user_id"`
	PlayURL       string    `gorm:"column:play_url;not null" json:"play_url"`             // 视频播放地址
	CoverURL      string    `gorm:"column:cover_url;not null" json:"cover_url"`           // 视频封面地址
	FavoriteCount int32     `gorm:"column:favorite_count;not null" json:"favorite_count"` // 视频点赞数
	CommentCount  int32     `gorm:"column:comment_count;not null" json:"comment_count"`   // 视频评论数
	Title         string    `gorm:"column:title;not null" json:"title"`
	CreateAt      time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt      time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP" json:"update_at"`
	Author        User      `gorm:"foreignKey:UserID" json:"author"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
