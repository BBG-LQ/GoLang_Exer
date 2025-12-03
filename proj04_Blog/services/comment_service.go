package services

import (
	"BLOG/models"
	"BLOG/req"
	"BLOG/utils"
	"time"
)

type CommentService struct {
}

func (s CommentService) CreateComment(request req.CreateComment) bool {

	var comment models.Comments

	comment.Content = request.Content
	comment.UserId = request.UserId
	comment.PostId = request.PostId
	comment.CreatedAt = time.Now().Truncate(time.Millisecond)
	comment.UpdatedAt = time.Now().Truncate(time.Millisecond)
	db := utils.DBUtil{}.Connect()

	db.Create(&comment)

	return true

}

func (s CommentService) GetCommentByPostId(request req.GetCommentByPostIdRequest) []models.Comments {
	var comments []models.Comments

	db := utils.DBUtil{}.Connect()
	db.Preload("Author").Preload("Post").Preload("ToUser").Preload("Post.Author").Where("post_id = ?", request.PostId).Find(&comments)

	return comments

}
