package models

type Types struct {
	ID       uint   `json:"id"`
	TypeName string `json:"typeName" gorm:"unique"`
}
