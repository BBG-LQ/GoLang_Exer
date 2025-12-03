package controller

import (
	"BLOG/req"
	"BLOG/services"

	"github.com/gin-gonic/gin"
)

func CommentControllerInit(r *gin.Engine) {
	comment := r.Group("/comment")
	{
		comment.POST("/createcomment", AuthMiddleware(), CreateComment)
		comment.POST("/getcommentbypostid", AuthMiddleware(), GetCommentByPostId)
	}
}

// 创建评论
func CreateComment(c *gin.Context) {
	var request req.CreateComment

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}
	request.UserId = c.GetInt("userId")
	ok := services.CommentService{}.CreateComment(request)
	if ok {
		c.JSON(200, gin.H{
			"message": "create successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "create fail",
		})
	}

}

// 获取文章评论
func GetCommentByPostId(c *gin.Context) {
	var request req.GetCommentByPostIdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}
	comments := services.CommentService{}.GetCommentByPostId(request)

	c.JSON(200, gin.H{
		"message": "create successfully",
		"data":    comments,
	})

}
