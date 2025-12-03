package models

import "time"

type Comments struct {
	Id        int       `json:"id" gorm:"id"`
	Content   string    `json:"content" gorm:"content"`
	UserId    int       `json:"userId" gorm:"user_id"`
	Author    Users     `json:"author" gorm:"foreignkey:UserId"`
	PostId    int       `json:"postId" gorm:"post_id"`
	Post      Posts     `json:"post" gorm:"foreignkey:PostId"`
	CreatedAt time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updated_at"`
}
