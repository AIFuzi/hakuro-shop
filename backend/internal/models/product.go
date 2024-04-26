package models

type Product struct {
	ID          uint   `json:"id"`
	TypeID      uint   `json:"typeID"`
	Image       string `json:"image"`
	ProductName string `json:"productName"`
}
