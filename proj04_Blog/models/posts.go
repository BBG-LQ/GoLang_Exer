package models

import "time"

type Posts struct {
	Id        int       `json:"id" gorm:"id"`
	Title     string    `json:"title" gorm:"title"`
	Content   string    `json:"content" gorm:"content"`
	UserId    int       `json:"userId" gorm:"user_id"`
	Author    Users     `json:"author" gorm:"foreignkey:UserId"` // 添加用户关联
	CreatedAt time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updated_at"`
}

func (Posts) TableName() string {
	return "posts"
}
