package requests

import "time"

type UserRequest struct {
	Name     string    `json:"name" binding:"required"`
	TID      string    `json:"tid" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Role     string    `json:"role" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
}
