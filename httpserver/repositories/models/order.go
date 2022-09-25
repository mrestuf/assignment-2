package models

import "time"

type Order struct {
	ID            int `gorm:"primaryKey; autoIncrement"`
	Customer_Name string
	Ordered_At    time.Time
	Items         []Item `json:"items" gorm:"foreignKey:Order_ID"`
}

type Item struct {
	ID          int `gorm:"primaryKey; autoIncrement"`
	Item_Code   string
	Description string
	Quantity    int
	Order_ID    int
}
