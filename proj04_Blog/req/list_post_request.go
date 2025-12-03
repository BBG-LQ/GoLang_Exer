package req

type ListPostRequest struct {
	PageNo   int `json:"pageNo" binding:"required"`
	PageSize int `json:"pageSize" binding:"required"`
}
