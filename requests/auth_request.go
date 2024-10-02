package requests

type LoginRequest struct {
	TID      string `json:"tid" form:"text" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
