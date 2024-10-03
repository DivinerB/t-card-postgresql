package models

import "time"

type Feedback struct {
	ID          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	State       *string    `gorm:"default:pending" json:"state"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	Feeder      *User      `gorm:"foreignKey:UserID"`
	UserID      *uint
}
