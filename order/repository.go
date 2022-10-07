package order

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get() ([]Order, error)
	Find(OrderID int) (Order, error)
	Create(order Order) (Order, error)
	Update(order Order) (Order, error)
	Delete(order Order) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Get() ([]Order, error) {
	var orders []Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r repository) Find(OrderID int) (Order, error) {
	var order Order
	err := r.db.First(&order, OrderID).Error
	return order, err
}

func (r repository) Create(order Order) (Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r repository) Update(order Order) (Order, error) {
	err := r.db.Save(&order).Error
	return order, err
}

func (r repository) Delete(order Order) (Order, error) {
	err := r.db.Delete(&order).Error
	return order, err
}
