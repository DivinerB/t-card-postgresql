package models

import "time"

type Job struct {
	ID           *uint          `json:"id"`
	Title        *string        `json:"title"`
	Description  *string        `json:"description"`
	Geoposition  *string        `json:"geoposision"`
	Salary       *int           `json:"salary"`
	Money        *string        `json:"money"`
	CreatedAt    *time.Time     `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
	Owner        *User          `gorm:"foreignKey:UserID"`
	Stacks       []*Stack       `gorm:"many2many:job_stacks"`
	Applications []*Application `gorm:"foreignKey:UserID"`
	Bookmarker   []*User        `gorm:"many2many:bookmarks"`
	UserID       *uint
}
