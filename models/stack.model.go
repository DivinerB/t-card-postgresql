package models

type Stack struct {
	ID   *uint   `json:"id"`
	Name *string `json:"name"`
	Logo *string `json:"logo"`
	Jobs []*Job  `gorm:"many2many:job_stacks"`
}
