package services

import (
	"BLOG/models"
	"BLOG/req"
	"BLOG/utils"
	"time"
)

type PostSercive struct {
}

// CreatePost 创建帖子
func (s PostSercive) CreatePost(request req.CreatePostReq) bool {
	var post models.Posts
	post.Title = request.Title
	post.Content = request.Content
	post.UserId = request.UserId
	post.CreatedAt = time.Now().Truncate(time.Millisecond)
	post.UpdatedAt = time.Now().Truncate(time.Millisecond)

	db := utils.DBUtil{}.Connect()
	db.Create(&post)
	return true
}

// DetailPost 获取帖子详情
func (s PostSercive) DetailPost(request req.DetailPostRequest) models.Posts {
	var post models.Posts

	db := utils.DBUtil{}.Connect()
	db.Where("id=?", request.Id).First(&post)

	return post
}

// DeletePost 删除帖子
func (s PostSercive) DeletePost(request req.DetailPostRequest) bool {
	db := utils.DBUtil{}.Connect()

	var post models.Posts
	db.Where("id=?", request.Id).First(&post)
	if post.Id == 0 {
		return false
	}
	if post.UserId != request.UserId {
		return false
	}

	return db.Where("id=?", request.Id).Where("user_id = ?", request.UserId).Delete(&models.Posts{}).Error == nil
}

// ListPost 获取帖子列表
func (s PostSercive) ListPost(request req.ListPostRequest) []models.Posts {
	var posts []models.Posts
	db := utils.DBUtil{}.Connect()
	db.Limit(request.PageSize).Offset((request.PageNo - 1) * request.PageSize).Find(&posts)

	return posts
}

// UpdatePost 修改帖子
func (s PostSercive) UpdatePost(request req.UpdatePostRequest) bool {

	db := utils.DBUtil{}.Connect()
	var post models.Posts
	db.Where("id = ?", request.Id).Find(&post)

	if post.Id == 0 {
		return false
	}
	db.Model(&post).Update("title", request.Title).Update("content", request.Content).Update("updated_at", time.Now())
	return true
}
