package services

import (
	"assignment2/httpserver/controllers/params"
	"assignment2/httpserver/controllers/views"
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"
	"database/sql"
	"strings"

	"github.com/jinzhu/gorm"
)

type OrderSvc struct {
	repo repositories.OrderRepo
}

func NewOrderSvc(repo repositories.OrderRepo) *OrderSvc {
	return &OrderSvc{
		repo: repo,
	}
}

func (o *OrderSvc) GetAllOrders() *views.Response {
	orders, err := o.repo.GetOrders()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessFindAllResponse(parseModelToOrderGetAll(orders), "GET_ALL_ORDERS")
}

func (o *OrderSvc) CreateOrder(req *params.OrderCreateRequest) *views.Response {
	order := parseRequestOrderToModel(req)

	err := o.repo.CreateOrder(order)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflict(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(&req, "Create Order Success!")
}

func (o *OrderSvc) UpdateOrder(id int, req *params.OrderUpdateRequest) *views.Response {
	order := parseUpdateRequestOrderToModel(req)

	err := o.repo.UpdateOrder(id, order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.BadRequestError(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(order, "Update Order Success!")
}

func (o *OrderSvc) DeleteOrderbyID(id int) *views.Response {

	err := o.repo.DeleteOrderbyID(id, &models.Order{})
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessDeleteResponse(err, "Delete Success!")
}

func parseUpdateRequestOrderToModel(req *params.OrderUpdateRequest) *models.Order {
	i := parseUpdateRequestItemToModel(req)
	return &models.Order{
		ID:            req.ID,
		Customer_Name: req.Customer_Name,
		Ordered_At:    req.Ordered_At,
		Items:         *i,
	}
}

func parseUpdateRequestItemToModel(req *params.OrderUpdateRequest) *[]models.Item {
	var i []models.Item
	for _, item := range req.Items {
		x := models.Item{
			ID:          item.ID,
			Item_Code:   item.Item_Code,
			Description: item.Description,
			Quantity:    item.Quantity,
			// Order_ID:    item.Order_ID,
		}
		i = append(i, x)
	}
	return &i
}

func parseRequestOrderToModel(req *params.OrderCreateRequest) *models.Order {
	i := parseRequestItemToModel(req)
	return &models.Order{
		Customer_Name: req.Customer_Name,
		Ordered_At:    req.Ordered_At,
		Items:         *i,
	}
}

func parseRequestItemToModel(req *params.OrderCreateRequest) *[]models.Item {
	var i []models.Item
	for _, item := range req.Items {
		x := models.Item{
			Item_Code:   item.Item_Code,
			Description: item.Description,
			Quantity:    item.Quantity,
		}
		i = append(i, x)
	}
	return &i
}

func parseModelToOrderGetAll(mod *[]models.Order) *[]views.OrderGetAll {
	var o []views.OrderGetAll
	for _, st := range *mod {
		i := []views.ItemsGet{}
		for _, item := range st.Items {
			i = append(i, views.ItemsGet{
				// ID:          item.ID,
				Item_Code:   item.Item_Code,
				Description: item.Description,
				Quantity:    item.Quantity,
			})

		}
		o = append(o, views.OrderGetAll{
			// ID:            st.ID,
			Customer_Name: st.Customer_Name,
			Ordered_At:    st.Ordered_At,
			Items:         i,
		})
	}
	return &o
}
