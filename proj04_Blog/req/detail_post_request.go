package req

type DetailPostRequest struct {
	Id     int `json:"id" binding:"required"`
	UserId int `json:"userId"`
}
