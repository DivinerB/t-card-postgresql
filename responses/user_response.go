package responses

import "time"

type UserResponse struct {
	Name     *string    `json:"name"`
	TID      *string    `json:"tid"`
	Role     *string    `json:"role"`
	Password *string    `json:"password"`
	Birthday *time.Time `json:"birthday"`
}
