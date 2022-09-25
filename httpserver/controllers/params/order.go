package params

import (
	"time"
)

type OrderCreateRequest struct {
	Ordered_At    time.Time           `validate:"required"`
	Customer_Name string              `validate:"required"`
	Items         []ItemCreateRequest `validate:"required"`
}

type ItemCreateRequest struct {
	Item_Code   string `validate:"required"`
	Description string `validate:"required"`
	Quantity    int    `validate:"required"`
}

type OrderUpdateRequest struct {
	ID            int
	Ordered_At    time.Time
	Customer_Name string
	Items         []ItemUpdateRequest
}

type ItemUpdateRequest struct {
	ID          int
	Item_Code   string
	Description string
	Quantity    int
	// Order_ID    int
}
