package req

type CreateComment struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserId  int    `json:"userId"`
	PostId  int    `json:"postId" binding:"required"`
}
