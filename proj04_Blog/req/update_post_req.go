package req

type UpdatePostRequest struct {
	Id      int    `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserId  int    `json:"userId"`
}
