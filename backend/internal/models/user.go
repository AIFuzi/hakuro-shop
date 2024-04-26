package models

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	Password []byte `json:"-"`
	Role     string `json:"role"`
}
