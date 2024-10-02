package models

import "time"

type Application struct {
	ID        *uint      `gorm:"primary_key" json:"id"`
	State     *string    `gorm:"default:pending" json:"state"`
	Applicant *User      `gorm:"foreignKey:UserID"`
	Job       *Job       `gorm:"foreignKey:JobID"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UserID    *uint
	JobID     *uint
}
