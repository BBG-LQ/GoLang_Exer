package controller

import (
	"BLOG/req"
	"BLOG/services"

	"github.com/gin-gonic/gin"
)

func PostControllerInit(r *gin.Engine) {
	post := r.Group("/post")
	{
		post.POST("/createPost", AuthMiddleware(), CreatePost)
		post.POST("/list", AuthMiddleware(), ListPost)
		post.POST("/detailpost", AuthMiddleware(), DetailPost)
		post.DELETE("/deletepost", AuthMiddleware(), DeletePost)
		post.PUT("/updatepost", AuthMiddleware(), UpdatePost)

	}
}

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	var request req.CreatePostReq

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}
	request.UserId = c.GetInt("userId")

	ok := services.PostSercive{}.CreatePost(request)
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

// DetailPost 获取帖子详情
func DetailPost(c *gin.Context) {
	var request req.DetailPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}

	post := services.PostSercive{}.DetailPost(request)
	c.JSON(200, gin.H{
		"message": "detailpost duccessfully",
		"data":    post,
	})

}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	var request req.DetailPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}
	request.UserId = c.GetInt("userId")
	ok := services.PostSercive{}.DeletePost(request)
	if ok {
		c.JSON(200, gin.H{
			"message": "deletepost duccessfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "delete fail",
		})
	}

}

// ListPost 获取帖子列表
func ListPost(c *gin.Context) {
	var request req.ListPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}

	posts := services.PostSercive{}.ListPost(request)
	c.JSON(200, gin.H{
		"message": "listpost duccessfully",
		"data":    posts,
	})

}

// UpdatePost 修改帖子
func UpdatePost(c *gin.Context) {
	var request req.UpdatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}

	ok := services.PostSercive{}.UpdatePost(request)
	if ok {
		c.JSON(200, gin.H{
			"message": "update duccessfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "update fail",
		})
	}

}
