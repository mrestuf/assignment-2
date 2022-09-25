package views

import "time"

type OrderGetAll struct {
	// ID            int        `json:"orderID"`
	Ordered_At    time.Time  `json:"orderedAt"`
	Customer_Name string     `json:"customerName"`
	Items         []ItemsGet `json:"items"`
}

type ItemsGet struct {
	// ID          int    `json:"itemID"`
	Item_Code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
