package req

type CreatePostReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserId  int    `json:"userId"`
}
