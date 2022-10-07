package order

import "assignment2/item"

type Service interface {
	Get() ([]Order, error)
	Create(req OrderRequest) (Order, error)
	Update(OrderID int, req OrderRequest) (Order, error)
	Delete(OrderID int) (Order, error)
}

type service struct {
	repository     Repository
	itemRepository item.Repository
}

func NewService(repository Repository, itemRepository item.Repository) *service {
	return &service{repository, itemRepository}
}

func (s *service) Get() ([]Order, error) {
	orders, err := s.repository.Get()
	return orders, err
}

func (s *service) Create(req OrderRequest) (Order, error) {
	order := Order{
		OrderedAt:    req.OrderedAt,
		CustomerName: req.CustomerName,
	}

	NewOrder, err := s.repository.Create(order)
	return NewOrder, err
}

func (s *service) Update(OrderID int, req OrderRequest) (Order, error) {

	order, err := s.repository.Find(OrderID)

	order.CustomerName = req.CustomerName
	order.OrderedAt = req.OrderedAt

	newOrder, err := s.repository.Update(order)
	return newOrder, err
}

func (s *service) Delete(OrderID int) (Order, error) {

	order, err := s.repository.Find(OrderID)

	newOrder, err := s.repository.Delete(order)
	return newOrder, err
}
