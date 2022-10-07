package item

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get() ([]Item, error)
	GetByFK(OrderID int) ([]Item, error)
	Find(ItemID int) (Item, error)
	Create(item Item) (Item, error)
	Update(item Item) (Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Get() ([]Item, error) {
	var items []Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r repository) Find(ItemID int) (Item, error) {
	var item Item
	err := r.db.Find(&item, ItemID).Error
	return item, err
}

func (r repository) GetByFK(OrderID int) ([]Item, error) {
	var items []Item
	err := r.db.Where("order_id = ?", OrderID).Find(&items).Error
	return items, err
}

func (r repository) Create(item Item) (Item, error) {
	err := r.db.Create(&item).Error
	return item, err
}

func (r repository) Update(item Item) (Item, error) {
	err := r.db.Save(&item).Error
	return item, err
}
