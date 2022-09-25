package repositories

import "assignment2/httpserver/repositories/models"

type OrderRepo interface {
	GetOrders() (*[]models.Order, error)
	CreateOrder(order *models.Order) error
	UpdateOrder(id int, order *models.Order) error
	DeleteOrderbyID(id int, order *models.Order) error
}
