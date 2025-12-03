package models

type Users struct {
	Id       int    `json:"id" gorm:"id" `
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
}

func (Users) TableName() string {
	return "users"
}
