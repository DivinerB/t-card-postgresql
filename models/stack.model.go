package models

type Stack struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
	Logo *string `json:"logo"`
	Jobs []*Job  `gorm:"many2many:job_stacks"`
}
