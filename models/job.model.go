package models

import "time"

type Job struct {
	ID          *int       `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Geoposition *string    `json:"geoposision"`
	Salary      *int       `json:"salary"`
	Money       *string    `json:"money"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	UserID      *int       `json:"user_id,omitempty"`
	Owner       *User      `gorm:"foreignKey:UserID"`
	Stacks      []*Stack   `gorm:"many2many:job_stacks"`
}
