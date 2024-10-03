package models

type Bookmark struct {
	ID     *uint `gorm:"primary_key" json:"id"`
	User   *User `gorm:"foreignKey:UserID"`
	Job    *Job  `gorm:"foreignKey:JobID"`
	UserID *uint
	JobID  *uint
}
