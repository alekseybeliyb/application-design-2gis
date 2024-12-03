package handler

import (
	"app-booking/internal/model"
)

type OrderRepository interface {
	CreateOrder(booking *model.Order) error
	GetAllOrders() ([]model.Order, error)
	Migrate() error
	ClearAndMigrate()
}
