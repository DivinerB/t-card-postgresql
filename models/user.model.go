package models

import "time"

type User struct {
	ID          *uint          `json:"id"`
	TID         *string        `json:"tid"`
	Name        *string        `json:"name"`
	Password    *string        `json:"password"`
	Role        *string        `json:"role"`
	Birthday    *time.Time     `json:"birthday"`
	Jobs        []*Job         `gorm:"foreignKey:UserID"`
	Application []*Application `gorm:"foriegnKey:UserID"`
}
