package gorm

import (
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"
	"database/sql"

	"github.com/jinzhu/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repositories.OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) GetOrders() (*[]models.Order, error) {
	var orders []models.Order
	err := o.db.Preload("Items").Find(&orders).Error

	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}
	return &orders, nil
}

func (o *orderRepo) CreateOrder(order *models.Order) error {
	err := o.db.Create(order).Error
	return err
}

func (o *orderRepo) UpdateOrder(id int, order *models.Order) error {
	if len(order.Items) >= 1 {
		err := o.db.Debug().Model(&order).Association("Items").Replace(order.Items).Error
		if err != nil {
			return err
		}
	}
	err := o.db.Where("id = ?", id).Updates(order).Error
	return err
}

func (o *orderRepo) DeleteOrderbyID(id int, order *models.Order) error {
	err := o.db.Where("id = ?", id).Delete(order).Error
	return err
}
