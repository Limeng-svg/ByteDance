// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRelation = "relation"

// Relation mapped from table <relation>
type Relation struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID     int32     `gorm:"column:user_id;not null" json:"user_id"`       // 关注者id
	ToUserID   int32     `gorm:"column:to_user_id;not null" json:"to_user_id"` // 被关注者id
	CreateAt   time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP" json:"create_at"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	FollowUser User      `gorm:"foreignKey:ToUserID" json:"follow_user"`
}

// TableName Relation's table name
func (*Relation) TableName() string {
	return TableNameRelation
}